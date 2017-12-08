package workflow

import (
	"errors"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/3dsim/auth0/auth0fakes"
	"github.com/3dsim/workflow-goclient/models"
	"github.com/go-openapi/swag"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

const (
	gatewayURL = "apiGatewayURL"
	workflowAPIBasePath = "workflow-api"
	audience = "test audience"
)

func TestNewClientExpectsClientReturned(t *testing.T) {
	// arrange
	// act
	client := NewClient(nil, gatewayURL, workflowAPIBasePath, audience)

	// assert
	assert.NotNil(t, client, "Expected new client to not be nil");
}

func TestWorkflow(t *testing.T) {
	// arrange
	workflowID := "my-workflow"
	endpoint := "/" + workflowAPIBasePath + "/workflows/{workflowID}"

	t.Run("WhenSuccessfulExpectsWorkflowReturned", func(t *testing.T) {
		// arrange
		fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
		fakeTokenFetcher.TokenReturns("token", nil)
		workflowToReturn := &models.Workflow{
			ID:   workflowID,
			State: "Running",
			WaitingOnCapacity: false,
		}
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
			receivedWorkflowID := mux.Vars(r)["workflowID"]
			assert.EqualValues(t, workflowID, receivedWorkflowID, "Expected workflow id received to match what was passed in")
			bytes, err := json.Marshal(workflowToReturn)
			if err != nil {
				t.Error("Failed to marshal workflow")
			}
			w.Write(bytes)
		})

		// Setup routes
		r := mux.NewRouter()
		r.HandleFunc(endpoint, handler)
		testServer := httptest.NewServer(r)
		defer testServer.Close()
		client := NewClient(fakeTokenFetcher, testServer.URL, workflowAPIBasePath, audience)

		// act
		workflow, err := client.Workflow(workflowID)

		// assert
		assert.Nil(t, err, "Expected error to be nil when getting workflow")
		assert.NotNil(t, workflow, "Expected retrieved workflow to not be nil")
		assert.Equal(t, workflowToReturn.ID, workflow.ID, "Expected workflow IDs to match")
		assert.Equal(t, workflowToReturn.State, workflow.State, "Expected workflow states to match")
		assert.Equal(t, workflowToReturn.WaitingOnCapacity, workflow.WaitingOnCapacity, "Expected workflow 'waiting on capacity' to match")
	})

	t.Run("WhenFetcherErrorsExpectsErrorReturned", func(t *testing.T) {
		// arrange
		expectedError := errors.New("Some auth0 error")
		fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
		fakeTokenFetcher.TokenReturns("", expectedError)
		client := NewClient(fakeTokenFetcher, gatewayURL, workflowAPIBasePath, audience)

		// act
		workflow, err := client.Workflow(workflowID)

		// assert
		assert.Nil(t, workflow, "Expected no workflow to be returned due to token error")
		assert.Equal(t, expectedError, err, "Expected an error returned")

	})

	t.Run("WhenAPIErrorsExpectsErrorReturned", func(t *testing.T) {
		// arrange
		fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
		fakeTokenFetcher.TokenReturns("Token", nil)

		// return server error from http handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
		})

		// set up routes
		r := mux.NewRouter()
		r.HandleFunc(endpoint, handler)
		testServer := httptest.NewServer(r)
		defer testServer.Close()
		client := NewClient(fakeTokenFetcher, testServer.URL, workflowAPIBasePath, audience)

		// act
		workflow, err := client.Workflow(workflowID)

		// assert
		assert.Nil(t, workflow, "Expected no workflow to be returned due to API error")
		assert.NotNil(t, err, "Expected an error returned because workflow API sent a 500 error")
	})
}

func TestCancelWorkflow(t *testing.T) {
	// arrange
	workflowID := "my-workflow"
	endpoint := "/" + workflowAPIBasePath + "/workflows/{workflowID}/cancel"

	t.Run("WhenSuccessfulExpectsNothingReturned", func(t *testing.T) {
		// arrange
		fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
		fakeTokenFetcher.TokenReturns("token", nil)

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
			receivedWorkflowID := mux.Vars(r)["workflowID"]
			assert.EqualValues(t, workflowID, receivedWorkflowID, "Expected workflow id received to match what was passed in")
		})

		// Setup routes
		r := mux.NewRouter()
		r.HandleFunc(endpoint, handler)
		testServer := httptest.NewServer(r)
		defer testServer.Close()
		client := NewClient(fakeTokenFetcher, testServer.URL, workflowAPIBasePath, audience)

		// act
		err := client.CancelWorkflow(workflowID)

		// assert
		assert.Nil(t, err, "Expected error to be nil when cancelling workflow")
	})

	t.Run("WhenFetcherErrorsExpectsErrorReturned", func(t *testing.T) {
		// arrange
		expectedError := errors.New("Some auth0 error")
		fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
		fakeTokenFetcher.TokenReturns("", expectedError)
		client := NewClient(fakeTokenFetcher, gatewayURL, workflowAPIBasePath, audience)

		// act
		err := client.CancelWorkflow(workflowID)

		// assert
		assert.Equal(t, expectedError, err, "Expected an error returned")

	})

	t.Run("WhenAPIErrorsExpectsErrorReturned", func(t *testing.T) {
		// arrange
		fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
		fakeTokenFetcher.TokenReturns("Token", nil)

		// return server error from http handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
		})

		// set up routes
		r := mux.NewRouter()
		r.HandleFunc(endpoint, handler)
		testServer := httptest.NewServer(r)
		defer testServer.Close()
		client := NewClient(fakeTokenFetcher, testServer.URL, workflowAPIBasePath, audience)

		// act
		err := client.CancelWorkflow(workflowID)

		// assert
		assert.NotNil(t, err, "Expected an error returned because workflow API sent a 500 error")
	})
}

func TestSignalWorkflow(t *testing.T) {
	// arrange
	workflowID := "my-workflow"
	endpoint := "/" + workflowAPIBasePath + "/workflows/{workflowID}/signals"

	t.Run("WhenSuccessfulExpectsNothingReturned", func(t *testing.T) {
		// arrange
		fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
		fakeTokenFetcher.TokenReturns("token", nil)
		signal := &models.Signal {
			Name : swag.String("SignalName"),
			Input : "inputjson",
		}
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
			receivedWorkflowID := mux.Vars(r)["workflowID"]
			assert.EqualValues(t, workflowID, receivedWorkflowID, "Expected workflow id received to match what was passed in")
			bodyBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			var actualSignal models.Signal
			err = json.Unmarshal(bodyBytes, &actualSignal)
			if err != nil {
				t.Error("Failed to unmarshal signal")
			}
			assert.EqualValues(t, *signal, actualSignal, "Expected signal to be passed in body of request")
		})

		// Setup routes
		r := mux.NewRouter()
		r.HandleFunc(endpoint, handler)
		testServer := httptest.NewServer(r)
		defer testServer.Close()
		client := NewClient(fakeTokenFetcher, testServer.URL, workflowAPIBasePath, audience)

		// act
		err := client.SignalWorkflow(workflowID, signal)

		// assert
		assert.Nil(t, err, "Expected error to be nil when signalling workflow")
	})

	t.Run("WhenFetcherErrorsExpectsErrorReturned", func(t *testing.T) {
		// arrange
		expectedError := errors.New("Some auth0 error")
		fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
		fakeTokenFetcher.TokenReturns("", expectedError)
		client := NewClient(fakeTokenFetcher, gatewayURL, workflowAPIBasePath, audience)

		// act
		err := client.SignalWorkflow(workflowID, nil)

		// assert
		assert.Equal(t, expectedError, err, "Expected an error returned")

	})

	t.Run("WhenAPIErrorsExpectsErrorReturned", func(t *testing.T) {
		// arrange
		fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
		fakeTokenFetcher.TokenReturns("Token", nil)

		// return server error from http handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
		})

		// set up routes
		r := mux.NewRouter()
		r.HandleFunc(endpoint, handler)
		testServer := httptest.NewServer(r)
		defer testServer.Close()
		client := NewClient(fakeTokenFetcher, testServer.URL, workflowAPIBasePath, audience)

		// act
		err := client.SignalWorkflow(workflowID, &models.Signal{})

		// assert
		assert.NotNil(t, err, "Expected an error returned because workflow API sent a 500 error")
	})
}

func TestUpdateActivity(t *testing.T) {
	// arrange
	workflowID := "my-workflow"
	activityID := "my-activity"
	endpoint := "/" + workflowAPIBasePath + "/workflows/{workflowID}/activities/{activityID}"
	activityToReturn := &models.Activity {
		ID:   swag.String(activityID),
		Status: swag.String("Completed"),
		PercentComplete: 100,
		Result: "json-string",
		Error: &models.ActivityError {
			Reason : swag.String("reason"),
			Details: "details",
		},
	}

	t.Run("WhenSuccessfulExpectsActivityReturned", func(t *testing.T) {
		// arrange
		fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
		fakeTokenFetcher.TokenReturns("token", nil)
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
			receivedWorkflowID := mux.Vars(r)["workflowID"]
			receivedActivityID := mux.Vars(r)["activityID"]
			assert.EqualValues(t, workflowID, receivedWorkflowID, "Expected workflow id received to match what was passed in")
			assert.EqualValues(t, activityID, receivedActivityID, "Expected activity id received to match what was passed in")
			bytes, err := json.Marshal(activityToReturn)
			if err != nil {
				t.Error("Failed to marshal activity")
			}
			w.Write(bytes)
		})

		// Setup routes
		r := mux.NewRouter()
		r.HandleFunc(endpoint, handler)
		testServer := httptest.NewServer(r)
		defer testServer.Close()
		client := NewClient(fakeTokenFetcher, testServer.URL, workflowAPIBasePath, audience)

		// act
		activity, err := client.UpdateActivity(workflowID, activityToReturn)

		// assert
		assert.Nil(t, err, "Expected error to be nil when getting workflow")
		assert.NotNil(t, activity, "Expected retrieved activity to not be nil")
		assert.Equal(t, *activityToReturn.ID, *activity.ID, "Expected activity IDs to match")
		assert.Equal(t, *activityToReturn.Status, *activity.Status, "Expected activity status to match")
		assert.Equal(t, activityToReturn.PercentComplete, activity.PercentComplete, "Expected activity percent complete to match")
		assert.Equal(t, activityToReturn.Result, activity.Result, "Expected activity result to match")
		assert.NotNil(t, activity.Error, "Expected activity error to not be nil")
		assert.Equal(t, *activityToReturn.Error.Reason, *activity.Error.Reason, "Expected activity error reason to not be nil")
		assert.Equal(t, activityToReturn.Error.Details, activity.Error.Details, "Expected activity error details to not be nil")
	})

	t.Run("WhenFetcherErrorsExpectsErrorReturned", func(t *testing.T) {
		// arrange
		expectedError := errors.New("Some auth0 error")
		fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
		fakeTokenFetcher.TokenReturns("", expectedError)
		client := NewClient(fakeTokenFetcher, gatewayURL, workflowAPIBasePath, audience)

		// act
		activity, err := client.UpdateActivity(workflowID, activityToReturn)

		// assert
		assert.Nil(t, activity, "Expected no activity to be returned due to token error")
		assert.Equal(t, expectedError, err, "Expected an error returned")

	})

	t.Run("WhenAPIErrorsExpectsErrorReturned", func(t *testing.T) {
		// arrange
		fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
		fakeTokenFetcher.TokenReturns("Token", nil)

		// return server error from http handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
		})

		// set up routes
		r := mux.NewRouter()
		r.HandleFunc(endpoint, handler)
		testServer := httptest.NewServer(r)
		defer testServer.Close()
		client := NewClient(fakeTokenFetcher, testServer.URL, workflowAPIBasePath, audience)

		// act
		activity, err := client.UpdateActivity(workflowID, activityToReturn)

		// assert
		assert.Nil(t, activity, "Expected no activity to be returned due to API error")
		assert.NotNil(t, err, "Expected an error returned because workflow API sent a 500 error")
	})
}

func TestHeartbeatActivity(t *testing.T) {
	// arrange
	workflowID := "my-workflow"
	activityID := "my-activity"
	taskToken := "token"
	heartbeatDetails := "details"
	endpoint := "/" + workflowAPIBasePath + "/workflows/{workflowID}/activities/{activityID}/heartbeat"
	heartbeatToReturn := &models.Heartbeat{
		ActivityID: swag.String(activityID),
		TaskToken: swag.String(taskToken),
		Cancelled: false,
		Details: heartbeatDetails,
	}

	t.Run("WhenSuccessfulExpectsHeartbeatReturned", func(t *testing.T) {
		// arrange
		fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
		fakeTokenFetcher.TokenReturns("token", nil)
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
			receivedWorkflowID := mux.Vars(r)["workflowID"]
			receivedActivityID := mux.Vars(r)["activityID"]
			assert.EqualValues(t, workflowID, receivedWorkflowID, "Expected workflow id received to match what was passed in")
			assert.EqualValues(t, activityID, receivedActivityID, "Expected activity id received to match what was passed in")
			bytes, err := json.Marshal(heartbeatToReturn)
			if err != nil {
				t.Error("Failed to marshal heartbeat")
			}
			w.Write(bytes)
		})

		// Setup routes
		r := mux.NewRouter()
		r.HandleFunc(endpoint, handler)
		testServer := httptest.NewServer(r)
		defer testServer.Close()
		client := NewClient(fakeTokenFetcher, testServer.URL, workflowAPIBasePath, audience)

		// act
		heartbeat, err := client.HeartbeatActivity(workflowID, activityID)

		// assert
		assert.Nil(t, err, "Expected error to be nil when getting workflow")
		assert.NotNil(t, heartbeat, "Expected retrieved heartbeat to not be nil")
		assert.Equal(t, *heartbeatToReturn.ActivityID, *heartbeat.ActivityID, "Expected heartbeat activity IDs to match")
		assert.Equal(t, *heartbeatToReturn.TaskToken, *heartbeat.TaskToken, "Expected heartbeat tokens to match")
		assert.Equal(t, heartbeatToReturn.Cancelled, heartbeat.Cancelled, "Expected heartbeat cancelled field to match")
		assert.Equal(t, heartbeatToReturn.Details, heartbeat.Details, "Expected heartbeat details to match")
	})

	t.Run("WhenFetcherErrorsExpectsErrorReturned", func(t *testing.T) {
		// arrange
		expectedError := errors.New("Some auth0 error")
		fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
		fakeTokenFetcher.TokenReturns("", expectedError)
		client := NewClient(fakeTokenFetcher, gatewayURL, workflowAPIBasePath, audience)

		// act
		heartbeat, err := client.HeartbeatActivity(workflowID, activityID)

		// assert
		assert.Nil(t, heartbeat, "Expected no heartbeat to be returned due to token error")
		assert.Equal(t, expectedError, err, "Expected an error returned")
	})

	t.Run("WhenAPIErrorsExpectsErrorReturned", func(t *testing.T) {
		// arrange
		fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
		fakeTokenFetcher.TokenReturns("Token", nil)

		// return server error from http handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
		})

		// set up routes
		r := mux.NewRouter()
		r.HandleFunc(endpoint, handler)
		testServer := httptest.NewServer(r)
		defer testServer.Close()
		client := NewClient(fakeTokenFetcher, testServer.URL, workflowAPIBasePath, audience)

		// act
		heartbeat, err := client.HeartbeatActivity(workflowID, activityID)

		// assert
		assert.Nil(t, heartbeat, "Expected no heartbeat to be returned due to API error")
		assert.NotNil(t, err, "Expected an error returned because workflow API sent a 500 error")
	})
}