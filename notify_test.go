package notify

import (
	"os"
	"testing"
)

func TestNotify_Send(t *testing.T) {
	type fields struct {
		config *Config
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"test pushover notify",
			fields{config: &Config{
				Platform: Platform("pushover"),
				Token:    os.Getenv("PUSHOVER_TOKEN"),
				Channel:  os.Getenv("PUSHOVER_CHANNEL"),
			}},
			args{msg: "test case"},
		},
		{
			"test slack notify",
			fields{config: &Config{
				Platform: Platform("slack"),
				Token:    os.Getenv("SLACK_TOKEN"),
				Channel:  os.Getenv("SLACK_CHANNEL"),
			}},
			args{msg: "test case"},
		},
		{
			"test pagerduty severity is null",
			fields{config: &Config{
				Platform: Platform("pagerduty"),
				Token:    os.Getenv("PAGERDUTY_TOKEN"),
				Source:   "api-test",
				Severity: "",
			}},
			args{msg: "test pagerduty"},
		},
		{
			"test pagerduty severity is error",
			fields{config: &Config{
				Platform: Platform("pagerduty"),
				Token:    os.Getenv("PAGERDUTY_TOKEN"),
				Source:   "api-test",
				Severity: "error",
			}},
			args{msg: "test pagerduty is error"},
		},
		{
			"test discord notify",
			fields{
				config: &Config{
					Platform: PlatformDiscord,
					Token:    os.Getenv("DISCORD_TOKEN"),
					Channel:  os.Getenv("DISCORD_CHANNEL"),
				},
			},
			args{msg: "test case"},
		},
		{
			name: "test telegram notify",
			fields: fields{
				config: &Config{
					Platform: PlatformTelegram,
					Token:    os.Getenv("TELEGRAM_TOKEN"),
					Channel:  os.Getenv("TELEGRAM_CHANNEL"),
				},
			},
			args: args{
				msg: "test case",
			},
		},
		{
			"test dingtalk notify",
			fields{config: &Config{
				Platform: PlatformDingTalk,
				Token:    os.Getenv("DingTalk_TOKEN"),
				Channel:  os.Getenv("DingTalk_CHANNEL"),
			}},
			args{msg: "test case"},
		},
		{
			"test email notify",
			fields{config: &Config{
				Platform: PlatformEmail,
				Token:    os.Getenv("Email_Token"),
				User:     os.Getenv("Email_User"),
				Password: os.Getenv("Email_Password"),
				Host:     os.Getenv("Email_Host"),
			}},
			args{msg: "test case"},
		},
		{
			"test ses notify",
			fields{config: &Config{
				Platform: PlatformSes,
				Token:    os.Getenv("Email_To"),
				Sender:   os.Getenv("Email_SENDER"),
				Key:      os.Getenv("IAM_KEY"),
				Secret:   os.Getenv("IAM_SECRET"),
				Area:     os.Getenv("IAM_AREA"),
			}},
			args{msg: "test case"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Notify{
				config: tt.fields.config,
			}
			err := n.Send(tt.args.msg)
			if err != nil {
				t.Errorf(err.Error())
			}
		})
	}
}
