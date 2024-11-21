package biz

import (
	"reflect"
	"testing"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

func TestCampaign_AddTask(t *testing.T) {
	type fields struct {
		Campaign model.Campaign
		Tasks    []*Task
	}
	type args struct {
		task *Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Campaign{
				Campaign: tt.fields.Campaign,
				Tasks:    tt.fields.Tasks,
			}
			if err := c.AddTask(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("AddTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCampaign_Complete(t *testing.T) {
	type fields struct {
		Campaign model.Campaign
		Tasks    []*Task
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Campaign{
				Campaign: tt.fields.Campaign,
				Tasks:    tt.fields.Tasks,
			}
			if err := c.Complete(); (err != nil) != tt.wantErr {
				t.Errorf("Complete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCampaign_Start(t *testing.T) {
	type fields struct {
		Campaign model.Campaign
		Tasks    []*Task
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Campaign{
				Campaign: tt.fields.Campaign,
				Tasks:    tt.fields.Tasks,
			}
			if err := c.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCampaign(t *testing.T) {
	type args struct {
		name    string
		startAt time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    *Campaign
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCampaign(tt.args.name, tt.args.startAt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCampaign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCampaign() got = %v, want %v", got, tt.want)
			}
		})
	}
}
