package gocd

/*
{
			"name": "stage1",
			"approved_by": "admin",
			"jobs": [
			  {
				"name": "job1",
				"result": "Failed",
				"state": "Completed",
				"id": 13,
				"scheduled_date": 1436172201081
			  }
			],
			"can_run": true,
			"result": "Failed",
			"approval_type": "success",
			"counter": "1",
			"id": 13,
			"operate_permission": true,
			"rerun_of_counter": null,
			"scheduled": true
		  }

  private String name;
    private long id;
    private JobHistory jobHistory;
    private boolean canRun;
    private boolean scheduled = true; // true if this stage history really happened
    private String approvalType;
    private String approvedBy;
    private String counter;
    private boolean operatePermission;
    private StageInstanceModel previousStage;
    private StageResult result;
    private StageIdentifier identifier;
	private Integer rerunOfCounter;
*/

// StageInstance represents the stage from the result from a pipeline run
type StageInstance struct {
	Name              string `json:"name"`
	ID                int    `json:"id"`
	Jobs              []*Job `json:"jobs,omitempty"`
	CanRun            bool   `json:"can_run"`
	Scheduled         bool   `json:"scheduled"`
	ApprovalType      string `json:"approval_type,omitempty"`
	ApprovedBy        string `json:"approved_by,omitempty"`
	Counter           string `json:"counter,omitempty"`
	OperatePermission bool   `json:"operate_permission,omitempty"`
	Result            string `json:"result,omitempty"`
	RerunOfCounter    string `json:"rerun_of_counter,omitempty"`
}
