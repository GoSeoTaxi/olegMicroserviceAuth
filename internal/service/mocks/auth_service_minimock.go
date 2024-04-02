// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/GoSeoTaxi/olegMicroserviceAuth/internal/service.AuthService -o auth_service_minimock.go -n AuthServiceMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/model"
	"github.com/gojuno/minimock/v3"
)

// AuthServiceMock implements service.AuthService
type AuthServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreate          func(ctx context.Context, info *model.User) (i1 int64, err error)
	inspectFuncCreate   func(ctx context.Context, info *model.User)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mAuthServiceMockCreate

	funcDelete          func(ctx context.Context, i1 int64) (err error)
	inspectFuncDelete   func(ctx context.Context, i1 int64)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mAuthServiceMockDelete

	funcGet          func(ctx context.Context, i1 int64) (up1 *model.User, err error)
	inspectFuncGet   func(ctx context.Context, i1 int64)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mAuthServiceMockGet

	funcUpdate          func(ctx context.Context, up1 *model.User) (err error)
	inspectFuncUpdate   func(ctx context.Context, up1 *model.User)
	afterUpdateCounter  uint64
	beforeUpdateCounter uint64
	UpdateMock          mAuthServiceMockUpdate
}

// NewAuthServiceMock returns a mock for service.AuthService
func NewAuthServiceMock(t minimock.Tester) *AuthServiceMock {
	m := &AuthServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mAuthServiceMockCreate{mock: m}
	m.CreateMock.callArgs = []*AuthServiceMockCreateParams{}

	m.DeleteMock = mAuthServiceMockDelete{mock: m}
	m.DeleteMock.callArgs = []*AuthServiceMockDeleteParams{}

	m.GetMock = mAuthServiceMockGet{mock: m}
	m.GetMock.callArgs = []*AuthServiceMockGetParams{}

	m.UpdateMock = mAuthServiceMockUpdate{mock: m}
	m.UpdateMock.callArgs = []*AuthServiceMockUpdateParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mAuthServiceMockCreate struct {
	mock               *AuthServiceMock
	defaultExpectation *AuthServiceMockCreateExpectation
	expectations       []*AuthServiceMockCreateExpectation

	callArgs []*AuthServiceMockCreateParams
	mutex    sync.RWMutex
}

// AuthServiceMockCreateExpectation specifies expectation struct of the AuthService.Create
type AuthServiceMockCreateExpectation struct {
	mock    *AuthServiceMock
	params  *AuthServiceMockCreateParams
	results *AuthServiceMockCreateResults
	Counter uint64
}

// AuthServiceMockCreateParams contains parameters of the AuthService.Create
type AuthServiceMockCreateParams struct {
	ctx  context.Context
	info *model.User
}

// AuthServiceMockCreateResults contains results of the AuthService.Create
type AuthServiceMockCreateResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for AuthService.Create
func (mmCreate *mAuthServiceMockCreate) Expect(ctx context.Context, info *model.User) *mAuthServiceMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("AuthServiceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &AuthServiceMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &AuthServiceMockCreateParams{ctx, info}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the AuthService.Create
func (mmCreate *mAuthServiceMockCreate) Inspect(f func(ctx context.Context, info *model.User)) *mAuthServiceMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for AuthServiceMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by AuthService.Create
func (mmCreate *mAuthServiceMockCreate) Return(i1 int64, err error) *AuthServiceMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("AuthServiceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &AuthServiceMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &AuthServiceMockCreateResults{i1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the AuthService.Create method
func (mmCreate *mAuthServiceMockCreate) Set(f func(ctx context.Context, info *model.User) (i1 int64, err error)) *AuthServiceMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the AuthService.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the AuthService.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the AuthService.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mAuthServiceMockCreate) When(ctx context.Context, info *model.User) *AuthServiceMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("AuthServiceMock.Create mock is already set by Set")
	}

	expectation := &AuthServiceMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &AuthServiceMockCreateParams{ctx, info},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up AuthService.Create return parameters for the expectation previously defined by the When method
func (e *AuthServiceMockCreateExpectation) Then(i1 int64, err error) *AuthServiceMock {
	e.results = &AuthServiceMockCreateResults{i1, err}
	return e.mock
}

// Create implements service.AuthService
func (mmCreate *AuthServiceMock) Create(ctx context.Context, info *model.User) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, info)
	}

	mm_params := AuthServiceMockCreateParams{ctx, info}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, &mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := AuthServiceMockCreateParams{ctx, info}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("AuthServiceMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the AuthServiceMock.Create")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, info)
	}
	mmCreate.t.Fatalf("Unexpected call to AuthServiceMock.Create. %v %v", ctx, info)
	return
}

// CreateAfterCounter returns a count of finished AuthServiceMock.Create invocations
func (mmCreate *AuthServiceMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of AuthServiceMock.Create invocations
func (mmCreate *AuthServiceMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to AuthServiceMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mAuthServiceMockCreate) Calls() []*AuthServiceMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*AuthServiceMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *AuthServiceMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *AuthServiceMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to AuthServiceMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to AuthServiceMock.Create")
		} else {
			m.t.Errorf("Expected call to AuthServiceMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to AuthServiceMock.Create")
	}
}

type mAuthServiceMockDelete struct {
	mock               *AuthServiceMock
	defaultExpectation *AuthServiceMockDeleteExpectation
	expectations       []*AuthServiceMockDeleteExpectation

	callArgs []*AuthServiceMockDeleteParams
	mutex    sync.RWMutex
}

// AuthServiceMockDeleteExpectation specifies expectation struct of the AuthService.Delete
type AuthServiceMockDeleteExpectation struct {
	mock    *AuthServiceMock
	params  *AuthServiceMockDeleteParams
	results *AuthServiceMockDeleteResults
	Counter uint64
}

// AuthServiceMockDeleteParams contains parameters of the AuthService.Delete
type AuthServiceMockDeleteParams struct {
	ctx context.Context
	i1  int64
}

// AuthServiceMockDeleteResults contains results of the AuthService.Delete
type AuthServiceMockDeleteResults struct {
	err error
}

// Expect sets up expected params for AuthService.Delete
func (mmDelete *mAuthServiceMockDelete) Expect(ctx context.Context, i1 int64) *mAuthServiceMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("AuthServiceMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &AuthServiceMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &AuthServiceMockDeleteParams{ctx, i1}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the AuthService.Delete
func (mmDelete *mAuthServiceMockDelete) Inspect(f func(ctx context.Context, i1 int64)) *mAuthServiceMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for AuthServiceMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by AuthService.Delete
func (mmDelete *mAuthServiceMockDelete) Return(err error) *AuthServiceMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("AuthServiceMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &AuthServiceMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &AuthServiceMockDeleteResults{err}
	return mmDelete.mock
}

// Set uses given function f to mock the AuthService.Delete method
func (mmDelete *mAuthServiceMockDelete) Set(f func(ctx context.Context, i1 int64) (err error)) *AuthServiceMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the AuthService.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the AuthService.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the AuthService.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mAuthServiceMockDelete) When(ctx context.Context, i1 int64) *AuthServiceMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("AuthServiceMock.Delete mock is already set by Set")
	}

	expectation := &AuthServiceMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &AuthServiceMockDeleteParams{ctx, i1},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up AuthService.Delete return parameters for the expectation previously defined by the When method
func (e *AuthServiceMockDeleteExpectation) Then(err error) *AuthServiceMock {
	e.results = &AuthServiceMockDeleteResults{err}
	return e.mock
}

// Delete implements service.AuthService
func (mmDelete *AuthServiceMock) Delete(ctx context.Context, i1 int64) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, i1)
	}

	mm_params := AuthServiceMockDeleteParams{ctx, i1}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, &mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := AuthServiceMockDeleteParams{ctx, i1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("AuthServiceMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the AuthServiceMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, i1)
	}
	mmDelete.t.Fatalf("Unexpected call to AuthServiceMock.Delete. %v %v", ctx, i1)
	return
}

// DeleteAfterCounter returns a count of finished AuthServiceMock.Delete invocations
func (mmDelete *AuthServiceMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of AuthServiceMock.Delete invocations
func (mmDelete *AuthServiceMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to AuthServiceMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mAuthServiceMockDelete) Calls() []*AuthServiceMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*AuthServiceMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *AuthServiceMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *AuthServiceMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to AuthServiceMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to AuthServiceMock.Delete")
		} else {
			m.t.Errorf("Expected call to AuthServiceMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to AuthServiceMock.Delete")
	}
}

type mAuthServiceMockGet struct {
	mock               *AuthServiceMock
	defaultExpectation *AuthServiceMockGetExpectation
	expectations       []*AuthServiceMockGetExpectation

	callArgs []*AuthServiceMockGetParams
	mutex    sync.RWMutex
}

// AuthServiceMockGetExpectation specifies expectation struct of the AuthService.Get
type AuthServiceMockGetExpectation struct {
	mock    *AuthServiceMock
	params  *AuthServiceMockGetParams
	results *AuthServiceMockGetResults
	Counter uint64
}

// AuthServiceMockGetParams contains parameters of the AuthService.Get
type AuthServiceMockGetParams struct {
	ctx context.Context
	i1  int64
}

// AuthServiceMockGetResults contains results of the AuthService.Get
type AuthServiceMockGetResults struct {
	up1 *model.User
	err error
}

// Expect sets up expected params for AuthService.Get
func (mmGet *mAuthServiceMockGet) Expect(ctx context.Context, i1 int64) *mAuthServiceMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("AuthServiceMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &AuthServiceMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &AuthServiceMockGetParams{ctx, i1}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the AuthService.Get
func (mmGet *mAuthServiceMockGet) Inspect(f func(ctx context.Context, i1 int64)) *mAuthServiceMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for AuthServiceMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by AuthService.Get
func (mmGet *mAuthServiceMockGet) Return(up1 *model.User, err error) *AuthServiceMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("AuthServiceMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &AuthServiceMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &AuthServiceMockGetResults{up1, err}
	return mmGet.mock
}

// Set uses given function f to mock the AuthService.Get method
func (mmGet *mAuthServiceMockGet) Set(f func(ctx context.Context, i1 int64) (up1 *model.User, err error)) *AuthServiceMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the AuthService.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the AuthService.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the AuthService.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mAuthServiceMockGet) When(ctx context.Context, i1 int64) *AuthServiceMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("AuthServiceMock.Get mock is already set by Set")
	}

	expectation := &AuthServiceMockGetExpectation{
		mock:   mmGet.mock,
		params: &AuthServiceMockGetParams{ctx, i1},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up AuthService.Get return parameters for the expectation previously defined by the When method
func (e *AuthServiceMockGetExpectation) Then(up1 *model.User, err error) *AuthServiceMock {
	e.results = &AuthServiceMockGetResults{up1, err}
	return e.mock
}

// Get implements service.AuthService
func (mmGet *AuthServiceMock) Get(ctx context.Context, i1 int64) (up1 *model.User, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, i1)
	}

	mm_params := AuthServiceMockGetParams{ctx, i1}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, &mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := AuthServiceMockGetParams{ctx, i1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("AuthServiceMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the AuthServiceMock.Get")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, i1)
	}
	mmGet.t.Fatalf("Unexpected call to AuthServiceMock.Get. %v %v", ctx, i1)
	return
}

// GetAfterCounter returns a count of finished AuthServiceMock.Get invocations
func (mmGet *AuthServiceMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of AuthServiceMock.Get invocations
func (mmGet *AuthServiceMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to AuthServiceMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mAuthServiceMockGet) Calls() []*AuthServiceMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*AuthServiceMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *AuthServiceMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *AuthServiceMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to AuthServiceMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to AuthServiceMock.Get")
		} else {
			m.t.Errorf("Expected call to AuthServiceMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to AuthServiceMock.Get")
	}
}

type mAuthServiceMockUpdate struct {
	mock               *AuthServiceMock
	defaultExpectation *AuthServiceMockUpdateExpectation
	expectations       []*AuthServiceMockUpdateExpectation

	callArgs []*AuthServiceMockUpdateParams
	mutex    sync.RWMutex
}

// AuthServiceMockUpdateExpectation specifies expectation struct of the AuthService.Update
type AuthServiceMockUpdateExpectation struct {
	mock    *AuthServiceMock
	params  *AuthServiceMockUpdateParams
	results *AuthServiceMockUpdateResults
	Counter uint64
}

// AuthServiceMockUpdateParams contains parameters of the AuthService.Update
type AuthServiceMockUpdateParams struct {
	ctx context.Context
	up1 *model.User
}

// AuthServiceMockUpdateResults contains results of the AuthService.Update
type AuthServiceMockUpdateResults struct {
	err error
}

// Expect sets up expected params for AuthService.Update
func (mmUpdate *mAuthServiceMockUpdate) Expect(ctx context.Context, up1 *model.User) *mAuthServiceMockUpdate {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("AuthServiceMock.Update mock is already set by Set")
	}

	if mmUpdate.defaultExpectation == nil {
		mmUpdate.defaultExpectation = &AuthServiceMockUpdateExpectation{}
	}

	mmUpdate.defaultExpectation.params = &AuthServiceMockUpdateParams{ctx, up1}
	for _, e := range mmUpdate.expectations {
		if minimock.Equal(e.params, mmUpdate.defaultExpectation.params) {
			mmUpdate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdate.defaultExpectation.params)
		}
	}

	return mmUpdate
}

// Inspect accepts an inspector function that has same arguments as the AuthService.Update
func (mmUpdate *mAuthServiceMockUpdate) Inspect(f func(ctx context.Context, up1 *model.User)) *mAuthServiceMockUpdate {
	if mmUpdate.mock.inspectFuncUpdate != nil {
		mmUpdate.mock.t.Fatalf("Inspect function is already set for AuthServiceMock.Update")
	}

	mmUpdate.mock.inspectFuncUpdate = f

	return mmUpdate
}

// Return sets up results that will be returned by AuthService.Update
func (mmUpdate *mAuthServiceMockUpdate) Return(err error) *AuthServiceMock {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("AuthServiceMock.Update mock is already set by Set")
	}

	if mmUpdate.defaultExpectation == nil {
		mmUpdate.defaultExpectation = &AuthServiceMockUpdateExpectation{mock: mmUpdate.mock}
	}
	mmUpdate.defaultExpectation.results = &AuthServiceMockUpdateResults{err}
	return mmUpdate.mock
}

// Set uses given function f to mock the AuthService.Update method
func (mmUpdate *mAuthServiceMockUpdate) Set(f func(ctx context.Context, up1 *model.User) (err error)) *AuthServiceMock {
	if mmUpdate.defaultExpectation != nil {
		mmUpdate.mock.t.Fatalf("Default expectation is already set for the AuthService.Update method")
	}

	if len(mmUpdate.expectations) > 0 {
		mmUpdate.mock.t.Fatalf("Some expectations are already set for the AuthService.Update method")
	}

	mmUpdate.mock.funcUpdate = f
	return mmUpdate.mock
}

// When sets expectation for the AuthService.Update which will trigger the result defined by the following
// Then helper
func (mmUpdate *mAuthServiceMockUpdate) When(ctx context.Context, up1 *model.User) *AuthServiceMockUpdateExpectation {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("AuthServiceMock.Update mock is already set by Set")
	}

	expectation := &AuthServiceMockUpdateExpectation{
		mock:   mmUpdate.mock,
		params: &AuthServiceMockUpdateParams{ctx, up1},
	}
	mmUpdate.expectations = append(mmUpdate.expectations, expectation)
	return expectation
}

// Then sets up AuthService.Update return parameters for the expectation previously defined by the When method
func (e *AuthServiceMockUpdateExpectation) Then(err error) *AuthServiceMock {
	e.results = &AuthServiceMockUpdateResults{err}
	return e.mock
}

// Update implements service.AuthService
func (mmUpdate *AuthServiceMock) Update(ctx context.Context, up1 *model.User) (err error) {
	mm_atomic.AddUint64(&mmUpdate.beforeUpdateCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdate.afterUpdateCounter, 1)

	if mmUpdate.inspectFuncUpdate != nil {
		mmUpdate.inspectFuncUpdate(ctx, up1)
	}

	mm_params := AuthServiceMockUpdateParams{ctx, up1}

	// Record call args
	mmUpdate.UpdateMock.mutex.Lock()
	mmUpdate.UpdateMock.callArgs = append(mmUpdate.UpdateMock.callArgs, &mm_params)
	mmUpdate.UpdateMock.mutex.Unlock()

	for _, e := range mmUpdate.UpdateMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmUpdate.UpdateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdate.UpdateMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdate.UpdateMock.defaultExpectation.params
		mm_got := AuthServiceMockUpdateParams{ctx, up1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdate.t.Errorf("AuthServiceMock.Update got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdate.UpdateMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdate.t.Fatal("No results are set for the AuthServiceMock.Update")
		}
		return (*mm_results).err
	}
	if mmUpdate.funcUpdate != nil {
		return mmUpdate.funcUpdate(ctx, up1)
	}
	mmUpdate.t.Fatalf("Unexpected call to AuthServiceMock.Update. %v %v", ctx, up1)
	return
}

// UpdateAfterCounter returns a count of finished AuthServiceMock.Update invocations
func (mmUpdate *AuthServiceMock) UpdateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdate.afterUpdateCounter)
}

// UpdateBeforeCounter returns a count of AuthServiceMock.Update invocations
func (mmUpdate *AuthServiceMock) UpdateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdate.beforeUpdateCounter)
}

// Calls returns a list of arguments used in each call to AuthServiceMock.Update.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdate *mAuthServiceMockUpdate) Calls() []*AuthServiceMockUpdateParams {
	mmUpdate.mutex.RLock()

	argCopy := make([]*AuthServiceMockUpdateParams, len(mmUpdate.callArgs))
	copy(argCopy, mmUpdate.callArgs)

	mmUpdate.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateDone returns true if the count of the Update invocations corresponds
// the number of defined expectations
func (m *AuthServiceMock) MinimockUpdateDone() bool {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateInspect logs each unmet expectation
func (m *AuthServiceMock) MinimockUpdateInspect() {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to AuthServiceMock.Update with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		if m.UpdateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to AuthServiceMock.Update")
		} else {
			m.t.Errorf("Expected call to AuthServiceMock.Update with params: %#v", *m.UpdateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		m.t.Error("Expected call to AuthServiceMock.Update")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *AuthServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateInspect()

			m.MinimockDeleteInspect()

			m.MinimockGetInspect()

			m.MinimockUpdateInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *AuthServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *AuthServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockDeleteDone() &&
		m.MinimockGetDone() &&
		m.MinimockUpdateDone()
}
