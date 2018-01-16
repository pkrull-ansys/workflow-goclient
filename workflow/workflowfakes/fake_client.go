// This file was generated by counterfeiter
package workflowfakes

import (
	"sync"

	"github.com/3dsim/workflow-goclient/models"
	"github.com/3dsim/workflow-goclient/workflow"
)

type FakeClient struct {
	WorkflowStub        func(workflowID string) (*models.Workflow, error)
	workflowMutex       sync.RWMutex
	workflowArgsForCall []struct {
		workflowID string
	}
	workflowReturns struct {
		result1 *models.Workflow
		result2 error
	}
	workflowReturnsOnCall map[int]struct {
		result1 *models.Workflow
		result2 error
	}
	CancelWorkflowStub        func(workflowID string) error
	cancelWorkflowMutex       sync.RWMutex
	cancelWorkflowArgsForCall []struct {
		workflowID string
	}
	cancelWorkflowReturns struct {
		result1 error
	}
	cancelWorkflowReturnsOnCall map[int]struct {
		result1 error
	}
	SignalWorkflowStub        func(workflowID string, signal *models.Signal) error
	signalWorkflowMutex       sync.RWMutex
	signalWorkflowArgsForCall []struct {
		workflowID string
		signal     *models.Signal
	}
	signalWorkflowReturns struct {
		result1 error
	}
	signalWorkflowReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateActivityStub        func(workflowID string, activity *models.Activity) (*models.Activity, error)
	updateActivityMutex       sync.RWMutex
	updateActivityArgsForCall []struct {
		workflowID string
		activity   *models.Activity
	}
	updateActivityReturns struct {
		result1 *models.Activity
		result2 error
	}
	updateActivityReturnsOnCall map[int]struct {
		result1 *models.Activity
		result2 error
	}
	UpdateActivityPercentCompleteStub        func(workflowID, activityID string, percentComplete int) (*models.Activity, error)
	updateActivityPercentCompleteMutex       sync.RWMutex
	updateActivityPercentCompleteArgsForCall []struct {
		workflowID      string
		activityID      string
		percentComplete int
	}
	updateActivityPercentCompleteReturns struct {
		result1 *models.Activity
		result2 error
	}
	updateActivityPercentCompleteReturnsOnCall map[int]struct {
		result1 *models.Activity
		result2 error
	}
	CompleteSuccessfulActivityStub        func(workflowID, activityID string, result interface{}) (*models.Activity, error)
	completeSuccessfulActivityMutex       sync.RWMutex
	completeSuccessfulActivityArgsForCall []struct {
		workflowID string
		activityID string
		result     interface{}
	}
	completeSuccessfulActivityReturns struct {
		result1 *models.Activity
		result2 error
	}
	completeSuccessfulActivityReturnsOnCall map[int]struct {
		result1 *models.Activity
		result2 error
	}
	CompleteCancelledActivityStub        func(workflowID, activityID, details string) (*models.Activity, error)
	completeCancelledActivityMutex       sync.RWMutex
	completeCancelledActivityArgsForCall []struct {
		workflowID string
		activityID string
		details    string
	}
	completeCancelledActivityReturns struct {
		result1 *models.Activity
		result2 error
	}
	completeCancelledActivityReturnsOnCall map[int]struct {
		result1 *models.Activity
		result2 error
	}
	CompleteFailedActivityStub        func(workflowID, activityID, reason, details string) (*models.Activity, error)
	completeFailedActivityMutex       sync.RWMutex
	completeFailedActivityArgsForCall []struct {
		workflowID string
		activityID string
		reason     string
		details    string
	}
	completeFailedActivityReturns struct {
		result1 *models.Activity
		result2 error
	}
	completeFailedActivityReturnsOnCall map[int]struct {
		result1 *models.Activity
		result2 error
	}
	HeartbeatActivityStub        func(workflowID string, activityID string) (*models.Heartbeat, error)
	heartbeatActivityMutex       sync.RWMutex
	heartbeatActivityArgsForCall []struct {
		workflowID string
		activityID string
	}
	heartbeatActivityReturns struct {
		result1 *models.Heartbeat
		result2 error
	}
	heartbeatActivityReturnsOnCall map[int]struct {
		result1 *models.Heartbeat
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) Workflow(workflowID string) (*models.Workflow, error) {
	fake.workflowMutex.Lock()
	ret, specificReturn := fake.workflowReturnsOnCall[len(fake.workflowArgsForCall)]
	fake.workflowArgsForCall = append(fake.workflowArgsForCall, struct {
		workflowID string
	}{workflowID})
	fake.recordInvocation("Workflow", []interface{}{workflowID})
	fake.workflowMutex.Unlock()
	if fake.WorkflowStub != nil {
		return fake.WorkflowStub(workflowID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.workflowReturns.result1, fake.workflowReturns.result2
}

func (fake *FakeClient) WorkflowCallCount() int {
	fake.workflowMutex.RLock()
	defer fake.workflowMutex.RUnlock()
	return len(fake.workflowArgsForCall)
}

func (fake *FakeClient) WorkflowArgsForCall(i int) string {
	fake.workflowMutex.RLock()
	defer fake.workflowMutex.RUnlock()
	return fake.workflowArgsForCall[i].workflowID
}

func (fake *FakeClient) WorkflowReturns(result1 *models.Workflow, result2 error) {
	fake.WorkflowStub = nil
	fake.workflowReturns = struct {
		result1 *models.Workflow
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) WorkflowReturnsOnCall(i int, result1 *models.Workflow, result2 error) {
	fake.WorkflowStub = nil
	if fake.workflowReturnsOnCall == nil {
		fake.workflowReturnsOnCall = make(map[int]struct {
			result1 *models.Workflow
			result2 error
		})
	}
	fake.workflowReturnsOnCall[i] = struct {
		result1 *models.Workflow
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CancelWorkflow(workflowID string) error {
	fake.cancelWorkflowMutex.Lock()
	ret, specificReturn := fake.cancelWorkflowReturnsOnCall[len(fake.cancelWorkflowArgsForCall)]
	fake.cancelWorkflowArgsForCall = append(fake.cancelWorkflowArgsForCall, struct {
		workflowID string
	}{workflowID})
	fake.recordInvocation("CancelWorkflow", []interface{}{workflowID})
	fake.cancelWorkflowMutex.Unlock()
	if fake.CancelWorkflowStub != nil {
		return fake.CancelWorkflowStub(workflowID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cancelWorkflowReturns.result1
}

func (fake *FakeClient) CancelWorkflowCallCount() int {
	fake.cancelWorkflowMutex.RLock()
	defer fake.cancelWorkflowMutex.RUnlock()
	return len(fake.cancelWorkflowArgsForCall)
}

func (fake *FakeClient) CancelWorkflowArgsForCall(i int) string {
	fake.cancelWorkflowMutex.RLock()
	defer fake.cancelWorkflowMutex.RUnlock()
	return fake.cancelWorkflowArgsForCall[i].workflowID
}

func (fake *FakeClient) CancelWorkflowReturns(result1 error) {
	fake.CancelWorkflowStub = nil
	fake.cancelWorkflowReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CancelWorkflowReturnsOnCall(i int, result1 error) {
	fake.CancelWorkflowStub = nil
	if fake.cancelWorkflowReturnsOnCall == nil {
		fake.cancelWorkflowReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cancelWorkflowReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) SignalWorkflow(workflowID string, signal *models.Signal) error {
	fake.signalWorkflowMutex.Lock()
	ret, specificReturn := fake.signalWorkflowReturnsOnCall[len(fake.signalWorkflowArgsForCall)]
	fake.signalWorkflowArgsForCall = append(fake.signalWorkflowArgsForCall, struct {
		workflowID string
		signal     *models.Signal
	}{workflowID, signal})
	fake.recordInvocation("SignalWorkflow", []interface{}{workflowID, signal})
	fake.signalWorkflowMutex.Unlock()
	if fake.SignalWorkflowStub != nil {
		return fake.SignalWorkflowStub(workflowID, signal)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.signalWorkflowReturns.result1
}

func (fake *FakeClient) SignalWorkflowCallCount() int {
	fake.signalWorkflowMutex.RLock()
	defer fake.signalWorkflowMutex.RUnlock()
	return len(fake.signalWorkflowArgsForCall)
}

func (fake *FakeClient) SignalWorkflowArgsForCall(i int) (string, *models.Signal) {
	fake.signalWorkflowMutex.RLock()
	defer fake.signalWorkflowMutex.RUnlock()
	return fake.signalWorkflowArgsForCall[i].workflowID, fake.signalWorkflowArgsForCall[i].signal
}

func (fake *FakeClient) SignalWorkflowReturns(result1 error) {
	fake.SignalWorkflowStub = nil
	fake.signalWorkflowReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) SignalWorkflowReturnsOnCall(i int, result1 error) {
	fake.SignalWorkflowStub = nil
	if fake.signalWorkflowReturnsOnCall == nil {
		fake.signalWorkflowReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.signalWorkflowReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) UpdateActivity(workflowID string, activity *models.Activity) (*models.Activity, error) {
	fake.updateActivityMutex.Lock()
	ret, specificReturn := fake.updateActivityReturnsOnCall[len(fake.updateActivityArgsForCall)]
	fake.updateActivityArgsForCall = append(fake.updateActivityArgsForCall, struct {
		workflowID string
		activity   *models.Activity
	}{workflowID, activity})
	fake.recordInvocation("UpdateActivity", []interface{}{workflowID, activity})
	fake.updateActivityMutex.Unlock()
	if fake.UpdateActivityStub != nil {
		return fake.UpdateActivityStub(workflowID, activity)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateActivityReturns.result1, fake.updateActivityReturns.result2
}

func (fake *FakeClient) UpdateActivityCallCount() int {
	fake.updateActivityMutex.RLock()
	defer fake.updateActivityMutex.RUnlock()
	return len(fake.updateActivityArgsForCall)
}

func (fake *FakeClient) UpdateActivityArgsForCall(i int) (string, *models.Activity) {
	fake.updateActivityMutex.RLock()
	defer fake.updateActivityMutex.RUnlock()
	return fake.updateActivityArgsForCall[i].workflowID, fake.updateActivityArgsForCall[i].activity
}

func (fake *FakeClient) UpdateActivityReturns(result1 *models.Activity, result2 error) {
	fake.UpdateActivityStub = nil
	fake.updateActivityReturns = struct {
		result1 *models.Activity
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) UpdateActivityReturnsOnCall(i int, result1 *models.Activity, result2 error) {
	fake.UpdateActivityStub = nil
	if fake.updateActivityReturnsOnCall == nil {
		fake.updateActivityReturnsOnCall = make(map[int]struct {
			result1 *models.Activity
			result2 error
		})
	}
	fake.updateActivityReturnsOnCall[i] = struct {
		result1 *models.Activity
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) UpdateActivityPercentComplete(workflowID string, activityID string, percentComplete int) (*models.Activity, error) {
	fake.updateActivityPercentCompleteMutex.Lock()
	ret, specificReturn := fake.updateActivityPercentCompleteReturnsOnCall[len(fake.updateActivityPercentCompleteArgsForCall)]
	fake.updateActivityPercentCompleteArgsForCall = append(fake.updateActivityPercentCompleteArgsForCall, struct {
		workflowID      string
		activityID      string
		percentComplete int
	}{workflowID, activityID, percentComplete})
	fake.recordInvocation("UpdateActivityPercentComplete", []interface{}{workflowID, activityID, percentComplete})
	fake.updateActivityPercentCompleteMutex.Unlock()
	if fake.UpdateActivityPercentCompleteStub != nil {
		return fake.UpdateActivityPercentCompleteStub(workflowID, activityID, percentComplete)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateActivityPercentCompleteReturns.result1, fake.updateActivityPercentCompleteReturns.result2
}

func (fake *FakeClient) UpdateActivityPercentCompleteCallCount() int {
	fake.updateActivityPercentCompleteMutex.RLock()
	defer fake.updateActivityPercentCompleteMutex.RUnlock()
	return len(fake.updateActivityPercentCompleteArgsForCall)
}

func (fake *FakeClient) UpdateActivityPercentCompleteArgsForCall(i int) (string, string, int) {
	fake.updateActivityPercentCompleteMutex.RLock()
	defer fake.updateActivityPercentCompleteMutex.RUnlock()
	return fake.updateActivityPercentCompleteArgsForCall[i].workflowID, fake.updateActivityPercentCompleteArgsForCall[i].activityID, fake.updateActivityPercentCompleteArgsForCall[i].percentComplete
}

func (fake *FakeClient) UpdateActivityPercentCompleteReturns(result1 *models.Activity, result2 error) {
	fake.UpdateActivityPercentCompleteStub = nil
	fake.updateActivityPercentCompleteReturns = struct {
		result1 *models.Activity
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) UpdateActivityPercentCompleteReturnsOnCall(i int, result1 *models.Activity, result2 error) {
	fake.UpdateActivityPercentCompleteStub = nil
	if fake.updateActivityPercentCompleteReturnsOnCall == nil {
		fake.updateActivityPercentCompleteReturnsOnCall = make(map[int]struct {
			result1 *models.Activity
			result2 error
		})
	}
	fake.updateActivityPercentCompleteReturnsOnCall[i] = struct {
		result1 *models.Activity
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CompleteSuccessfulActivity(workflowID string, activityID string, result interface{}) (*models.Activity, error) {
	fake.completeSuccessfulActivityMutex.Lock()
	ret, specificReturn := fake.completeSuccessfulActivityReturnsOnCall[len(fake.completeSuccessfulActivityArgsForCall)]
	fake.completeSuccessfulActivityArgsForCall = append(fake.completeSuccessfulActivityArgsForCall, struct {
		workflowID string
		activityID string
		result     interface{}
	}{workflowID, activityID, result})
	fake.recordInvocation("CompleteSuccessfulActivity", []interface{}{workflowID, activityID, result})
	fake.completeSuccessfulActivityMutex.Unlock()
	if fake.CompleteSuccessfulActivityStub != nil {
		return fake.CompleteSuccessfulActivityStub(workflowID, activityID, result)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.completeSuccessfulActivityReturns.result1, fake.completeSuccessfulActivityReturns.result2
}

func (fake *FakeClient) CompleteSuccessfulActivityCallCount() int {
	fake.completeSuccessfulActivityMutex.RLock()
	defer fake.completeSuccessfulActivityMutex.RUnlock()
	return len(fake.completeSuccessfulActivityArgsForCall)
}

func (fake *FakeClient) CompleteSuccessfulActivityArgsForCall(i int) (string, string, interface{}) {
	fake.completeSuccessfulActivityMutex.RLock()
	defer fake.completeSuccessfulActivityMutex.RUnlock()
	return fake.completeSuccessfulActivityArgsForCall[i].workflowID, fake.completeSuccessfulActivityArgsForCall[i].activityID, fake.completeSuccessfulActivityArgsForCall[i].result
}

func (fake *FakeClient) CompleteSuccessfulActivityReturns(result1 *models.Activity, result2 error) {
	fake.CompleteSuccessfulActivityStub = nil
	fake.completeSuccessfulActivityReturns = struct {
		result1 *models.Activity
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CompleteSuccessfulActivityReturnsOnCall(i int, result1 *models.Activity, result2 error) {
	fake.CompleteSuccessfulActivityStub = nil
	if fake.completeSuccessfulActivityReturnsOnCall == nil {
		fake.completeSuccessfulActivityReturnsOnCall = make(map[int]struct {
			result1 *models.Activity
			result2 error
		})
	}
	fake.completeSuccessfulActivityReturnsOnCall[i] = struct {
		result1 *models.Activity
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CompleteCancelledActivity(workflowID string, activityID string, details string) (*models.Activity, error) {
	fake.completeCancelledActivityMutex.Lock()
	ret, specificReturn := fake.completeCancelledActivityReturnsOnCall[len(fake.completeCancelledActivityArgsForCall)]
	fake.completeCancelledActivityArgsForCall = append(fake.completeCancelledActivityArgsForCall, struct {
		workflowID string
		activityID string
		details    string
	}{workflowID, activityID, details})
	fake.recordInvocation("CompleteCancelledActivity", []interface{}{workflowID, activityID, details})
	fake.completeCancelledActivityMutex.Unlock()
	if fake.CompleteCancelledActivityStub != nil {
		return fake.CompleteCancelledActivityStub(workflowID, activityID, details)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.completeCancelledActivityReturns.result1, fake.completeCancelledActivityReturns.result2
}

func (fake *FakeClient) CompleteCancelledActivityCallCount() int {
	fake.completeCancelledActivityMutex.RLock()
	defer fake.completeCancelledActivityMutex.RUnlock()
	return len(fake.completeCancelledActivityArgsForCall)
}

func (fake *FakeClient) CompleteCancelledActivityArgsForCall(i int) (string, string, string) {
	fake.completeCancelledActivityMutex.RLock()
	defer fake.completeCancelledActivityMutex.RUnlock()
	return fake.completeCancelledActivityArgsForCall[i].workflowID, fake.completeCancelledActivityArgsForCall[i].activityID, fake.completeCancelledActivityArgsForCall[i].details
}

func (fake *FakeClient) CompleteCancelledActivityReturns(result1 *models.Activity, result2 error) {
	fake.CompleteCancelledActivityStub = nil
	fake.completeCancelledActivityReturns = struct {
		result1 *models.Activity
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CompleteCancelledActivityReturnsOnCall(i int, result1 *models.Activity, result2 error) {
	fake.CompleteCancelledActivityStub = nil
	if fake.completeCancelledActivityReturnsOnCall == nil {
		fake.completeCancelledActivityReturnsOnCall = make(map[int]struct {
			result1 *models.Activity
			result2 error
		})
	}
	fake.completeCancelledActivityReturnsOnCall[i] = struct {
		result1 *models.Activity
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CompleteFailedActivity(workflowID string, activityID string, reason string, details string) (*models.Activity, error) {
	fake.completeFailedActivityMutex.Lock()
	ret, specificReturn := fake.completeFailedActivityReturnsOnCall[len(fake.completeFailedActivityArgsForCall)]
	fake.completeFailedActivityArgsForCall = append(fake.completeFailedActivityArgsForCall, struct {
		workflowID string
		activityID string
		reason     string
		details    string
	}{workflowID, activityID, reason, details})
	fake.recordInvocation("CompleteFailedActivity", []interface{}{workflowID, activityID, reason, details})
	fake.completeFailedActivityMutex.Unlock()
	if fake.CompleteFailedActivityStub != nil {
		return fake.CompleteFailedActivityStub(workflowID, activityID, reason, details)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.completeFailedActivityReturns.result1, fake.completeFailedActivityReturns.result2
}

func (fake *FakeClient) CompleteFailedActivityCallCount() int {
	fake.completeFailedActivityMutex.RLock()
	defer fake.completeFailedActivityMutex.RUnlock()
	return len(fake.completeFailedActivityArgsForCall)
}

func (fake *FakeClient) CompleteFailedActivityArgsForCall(i int) (string, string, string, string) {
	fake.completeFailedActivityMutex.RLock()
	defer fake.completeFailedActivityMutex.RUnlock()
	return fake.completeFailedActivityArgsForCall[i].workflowID, fake.completeFailedActivityArgsForCall[i].activityID, fake.completeFailedActivityArgsForCall[i].reason, fake.completeFailedActivityArgsForCall[i].details
}

func (fake *FakeClient) CompleteFailedActivityReturns(result1 *models.Activity, result2 error) {
	fake.CompleteFailedActivityStub = nil
	fake.completeFailedActivityReturns = struct {
		result1 *models.Activity
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CompleteFailedActivityReturnsOnCall(i int, result1 *models.Activity, result2 error) {
	fake.CompleteFailedActivityStub = nil
	if fake.completeFailedActivityReturnsOnCall == nil {
		fake.completeFailedActivityReturnsOnCall = make(map[int]struct {
			result1 *models.Activity
			result2 error
		})
	}
	fake.completeFailedActivityReturnsOnCall[i] = struct {
		result1 *models.Activity
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) HeartbeatActivity(workflowID string, activityID string) (*models.Heartbeat, error) {
	fake.heartbeatActivityMutex.Lock()
	ret, specificReturn := fake.heartbeatActivityReturnsOnCall[len(fake.heartbeatActivityArgsForCall)]
	fake.heartbeatActivityArgsForCall = append(fake.heartbeatActivityArgsForCall, struct {
		workflowID string
		activityID string
	}{workflowID, activityID})
	fake.recordInvocation("HeartbeatActivity", []interface{}{workflowID, activityID})
	fake.heartbeatActivityMutex.Unlock()
	if fake.HeartbeatActivityStub != nil {
		return fake.HeartbeatActivityStub(workflowID, activityID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.heartbeatActivityReturns.result1, fake.heartbeatActivityReturns.result2
}

func (fake *FakeClient) HeartbeatActivityCallCount() int {
	fake.heartbeatActivityMutex.RLock()
	defer fake.heartbeatActivityMutex.RUnlock()
	return len(fake.heartbeatActivityArgsForCall)
}

func (fake *FakeClient) HeartbeatActivityArgsForCall(i int) (string, string) {
	fake.heartbeatActivityMutex.RLock()
	defer fake.heartbeatActivityMutex.RUnlock()
	return fake.heartbeatActivityArgsForCall[i].workflowID, fake.heartbeatActivityArgsForCall[i].activityID
}

func (fake *FakeClient) HeartbeatActivityReturns(result1 *models.Heartbeat, result2 error) {
	fake.HeartbeatActivityStub = nil
	fake.heartbeatActivityReturns = struct {
		result1 *models.Heartbeat
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) HeartbeatActivityReturnsOnCall(i int, result1 *models.Heartbeat, result2 error) {
	fake.HeartbeatActivityStub = nil
	if fake.heartbeatActivityReturnsOnCall == nil {
		fake.heartbeatActivityReturnsOnCall = make(map[int]struct {
			result1 *models.Heartbeat
			result2 error
		})
	}
	fake.heartbeatActivityReturnsOnCall[i] = struct {
		result1 *models.Heartbeat
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.workflowMutex.RLock()
	defer fake.workflowMutex.RUnlock()
	fake.cancelWorkflowMutex.RLock()
	defer fake.cancelWorkflowMutex.RUnlock()
	fake.signalWorkflowMutex.RLock()
	defer fake.signalWorkflowMutex.RUnlock()
	fake.updateActivityMutex.RLock()
	defer fake.updateActivityMutex.RUnlock()
	fake.updateActivityPercentCompleteMutex.RLock()
	defer fake.updateActivityPercentCompleteMutex.RUnlock()
	fake.completeSuccessfulActivityMutex.RLock()
	defer fake.completeSuccessfulActivityMutex.RUnlock()
	fake.completeCancelledActivityMutex.RLock()
	defer fake.completeCancelledActivityMutex.RUnlock()
	fake.completeFailedActivityMutex.RLock()
	defer fake.completeFailedActivityMutex.RUnlock()
	fake.heartbeatActivityMutex.RLock()
	defer fake.heartbeatActivityMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ workflow.Client = new(FakeClient)
