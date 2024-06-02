package process

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Arcadian-Sky/musthave-exam/internal/mocks"
	"github.com/Arcadian-Sky/musthave-exam/internal/model"
	"github.com/Arcadian-Sky/musthave-exam/internal/process/repolistener"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewOrderProcess(t *testing.T) {
	logger := logrus.New()
	repo := new(mocks.MockRepository)
	address := "http://example.com"

	process := NewOrderProcess(repo, logger, address)

	assert.Equal(t, 1*time.Second, process.iterTime)
	assert.Equal(t, 1*time.Minute, process.pauseTime)
	assert.Equal(t, repo, process.repo)
	assert.Equal(t, logger, process.log)
	assert.Equal(t, address, process.address)
	assert.NotNil(t, process.Listener)
}

func TestProcessTransaction(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer mockDB.Close()

	logger := logrus.New()

	repo := mocks.NewMockRepository(mockDB)

	address := "http://example.com"
	process := NewOrderProcess(repo, logger, address)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	process.Listener = repolistener.NewListener()

	// tx, err := mockDB.Begin()
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT id, status, action, date, summ FROM transactions WHERE id = ?`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "status", "action", "date", "summ"}).
			AddRow(1, "NEW", "Debit", time.Now(), 100))

	mock.ExpectExec(`UPDATE transactions SET status = ?, summ = ? WHERE id = ?`).
		WithArgs("PROCESSING", 100.0, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery(`SELECT user_id FROM transactions WHERE id = ?`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"user_id"}).
			AddRow(1))

	mock.ExpectExec(`UPDATE "user" SET balance = balance \+ ? WHERE id = ?`).
		WithArgs(100.0, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	// go func() {
	// 	process.processTransaction(ctx, model.Transaction{ID: 1, Status: "NEW"})
	// }()

	time.Sleep(100 * time.Millisecond)

	assert.NoError(t, mock.ExpectationsWereMet())

	// assert.Empty(t, hook.AllEntries())
}

func TestProcessTransaction_FetchAccrualError(t *testing.T) {
	// Создаем логгер и мок репозитория
	logger := logrus.New()
	mockRepo, mockDB, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	// Создаем репозиторий
	repo := repository.NewSQLRepository(mockDB)

	// Создаем процесс заказа
	address := "http://example.com"
	process := NewOrderProcess(repo, logger, address)

	// Создаем контекст и отменяем его после завершения теста
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создаем имитацию listener для новых заказов
	process.Listener = repolistener.NewListener()

	// Начинаем мок транзакции
	tx, err := mockDB.Begin()
	assert.NoError(t, err)

	// Ожидаемый вызов запроса на обновление статуса транзакции
	mockDB.ExpectBegin()
	mockDB.ExpectQuery(`SELECT id, status, action, date, summ FROM transactions WHERE id = ?`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "status", "action", "date", "summ"}).
			AddRow(1, "NEW", "Debit", time.Now(), 100))

	// Ожидаемый вызов обновления статуса
	mockDB.ExpectExec(`UPDATE transactions SET status = ?, summ = ? WHERE id = ?`).
		WithArgs("PROCESSING", 100.0, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Ожидаем ошибку при вызове FetchAccrual
	expectedError := errors.New("fetch accrual error")
	mockDB.ExpectQuery(`SELECT id FROM transactions WHERE id = ?`).
		WithArgs(1).
		WillReturnError(expectedError)

	// Ожидаем откат транзакции
	mockDB.ExpectRollback()

	// Запускаем процесс транзакции
	go func() {
		process.processTransaction(ctx, model.Transaction{ID: 1, Status: "NEW"})
	}()

	// Ждем выполнения ожиданий запросов
	time.Sleep(100 * time.Millisecond)

	// Проверяем, что все ожидаемые вызовы прошли успешно
	assert.NoError(t, mockDB.ExpectationsWereMet())

	// Проверяем, что в логах есть ошибка при вызове FetchAccrual
	assert.EqualError(t, hook.LastEntry().Cause, expectedError.Error())
}

func TestProcessTransaction_UpdateStatusError(t *testing.T) {
	// Создаем логгер и мок репозитория
	logger := logrus.New()
	mockRepo, mockDB, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	// Создаем репозиторий
	repo := repository.NewSQLRepository(mockDB)

	// Создаем процесс заказа
	address := "http://example.com"
	process := NewOrderProcess(repo, logger, address)

	// Создаем контекст и отменяем его после завершения теста
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создаем имитацию listener для новых заказов
	process.Listener = repolistener.NewListener()

	// Начинаем мок транзакции
	tx, err := mockDB.Begin()
	assert.NoError(t, err)

	// Ожидаемый вызов запроса на обновление статуса транзакции
	mockDB.ExpectBegin()
	mockDB.ExpectQuery(`SELECT id, status, action, date, summ FROM transactions WHERE id = ?`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "status", "action", "date", "summ"}).
			AddRow(1, "NEW", "Debit", time.Now(), 100))

	// Ожидаемый вызов обновления статуса
	mockDB.ExpectExec(`UPDATE transactions SET status = ?, summ = ? WHERE id = ?`).
		WithArgs("PROCESSING", 100.0, 1).
		WillReturnError(errors.New("update status error"))

	// Ожидаем откат транзакции
	mockDB.ExpectRollback()

	// Запускаем процесс транзакции
	go func() {
		process.processTransaction(ctx, model.Transaction{ID: 1, Status: "NEW"})
	}()

	// Ждем выполнения ожиданий запросов
	time.Sleep(100 * time.Millisecond)

	// Проверяем, что все ожидаемые вызовы прошли успешно
	assert.NoError(t, mockDB.ExpectationsWereMet())

	// Проверяем, что в логах есть ошибка при обновлении статуса
	assert.EqualError(t, hook.LastEntry().Cause, "update status error")
}

func TestWaitTransactionProcessing(t *testing.T) {
	// logger := logrus.New()
	// repo := new(mocks.MockRepository)
	// address := "http://example.com"

	// process := NewOrderProcess(repo, logger, address)

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// orderID := int64(1)
	// go func() {
	// 	time.Sleep(100 * time.Millisecond)
	// 	process.Listener.NewOrderChan <- orderID
	// 	cancel()
	// }()

	// process.WaitTransactionProcessing(ctx)
}

func TestStartTransactionProcessing(t *testing.T) {
	// logger := logrus.New()
	// repo := new(mocks.MockRepository)
	// address := "http://example.com"

	// process := NewOrderProcess(repo, logger, address)

	// ctx := context.Background()

	// mock, err := sqlmock.New()
	// assert.NoError(t, err)
	// defer mock.Close()

	// process.repo = repository.NewSQLRepository(mock)

	// mock.ExpectQuery(`SELECT id, user_id, summ, date, status, action FROM transactions WHERE action = 'Debit' AND status = 'NEW'`).
	// 	WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "summ", "date", "status", "action"}).
	// 		AddRow(1, 1, 100, time.Now(), "NEW", "Debit"))

	// process.StartTransactionProcessing(ctx)

	// assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProcessTransaction(t *testing.T) {
	// logger := logrus.New()
	// repo := new(mocks.MockRepository)
	// address := "http://example.com"

	// process := NewOrderProcess(repo, logger, address)

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// transaction := model.Transaction{ID: 1, Status: "NEW"}

	// go process.processTransaction(ctx, transaction)

	// time.Sleep(100 * time.Millisecond)
	// cancel()
}
