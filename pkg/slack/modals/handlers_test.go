package modals

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/slack-go/slack"
)

func TestValuesFor(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []byte
		blockIds []string
		expected map[string]string
	}{
		{
			name:     "complex input with selectors",
			input:    []byte(`{"action_id":"","action_ts":"","actions":[],"api_app_id":"A01BJF00CAD","attachment_id":"","block_id":"","callback_id":"","channel":{"created":0,"creator":"","id":"","is_archived":false,"is_channel":false,"is_ext_shared":false,"is_general":false,"is_group":false,"is_im":false,"is_member":false,"is_mpim":false,"is_open":false,"is_org_shared":false,"is_pending_ext_shared":false,"is_private":false,"is_shared":false,"locale":"","members":null,"name":"","name_normalized":"","num_members":0,"priority":0,"purpose":{"creator":"","last_set":0,"value":""},"topic":{"creator":"","last_set":0,"value":""},"unlinked":0,"user":""},"container":{"attachment_id":0,"channel_id":"","is_app_unfurl":false,"is_ephemeral":false,"message_ts":"","type":"","view_id":""},"hash":"","is_cleared":false,"message":{"blocks":null,"delete_original":false,"replace_original":false},"message_ts":"","name":"","original_message":{"blocks":null,"delete_original":false,"replace_original":false},"response_url":"","submission":null,"team":{"domain":"dptp-robot-testing","id":"T01B37EA9T5","name":""},"token":"Rn1sVT099UjUeKclB3PkLT8t","trigger_id":"1417512545924.1377252349923.130cb38151c96634bf8c75ffd162b735","type":"view_submission","user":{"color":"","deleted":false,"enterprise_user":{"enterprise_id":"","enterprise_name":"","id":"","is_admin":false,"is_owner":false,"teams":null},"has_2fa":false,"has_files":false,"id":"U01B31ARZDG","is_admin":false,"is_app_user":false,"is_bot":false,"is_invited_user":false,"is_owner":false,"is_primary_owner":false,"is_restricted":false,"is_stranger":false,"is_ultra_restricted":false,"locale":"","name":"skuznets","presence":"","profile":{"display_name":"","display_name_normalized":"","email":"","fields":[],"first_name":"","image_192":"","image_24":"","image_32":"","image_48":"","image_72":"","image_original":"","last_name":"","phone":"","real_name":"","real_name_normalized":"","skype":"","status_expiration":0,"team":"","title":""},"real_name":"","team_id":"T01B37EA9T5","tz_label":"","tz_offset":0,"updated":0},"value":"","view":{"app_id":"A01BJF00CAD","blocks":[{"block_id":"cANW","text":{"emoji":true,"text":"Members of the Test Platform team can use this form to document incidents and automatically create incident cards in Jira.","type":"plain_text"},"type":"section"},{"accessory":{"action_id":"EtZ","text":{"emoji":true,"text":"Triage an Incident","type":"plain_text"},"type":"button","value":"triage"},"block_id":"Cxauc","text":{"emoji":true,"text":"Users that wish to report an ongoing incident to engage the Test Platform Triage role should use the incident report form instead.","type":"plain_text"},"type":"section"},{"block_id":"Yv=","type":"divider"},{"block_id":"title","element":{"action_id":"vyRE","type":"plain_text_input"},"label":{"emoji":true,"text":"Provide a title for this incident:","type":"plain_text"},"type":"input"},{"block_id":"summary","element":{"action_id":"lEC","multiline":true,"type":"plain_text_input"},"label":{"emoji":true,"text":"Summarize what is happening:","type":"plain_text"},"type":"input"},{"block_id":"impact","element":{"action_id":"fKxhJ","multiline":true,"type":"plain_text_input"},"label":{"emoji":true,"text":"Explain the impact:","type":"plain_text"},"type":"input"},{"block_id":"bugzilla","element":{"action_id":"VejTS","type":"plain_text_input"},"label":{"emoji":true,"text":"Link the Bugzilla bug:","type":"plain_text"},"type":"input"},{"block_id":"selectors","elements":[{"action_id":"cJ/","placeholder":{"emoji":true,"text":"Select the incident channel...","type":"plain_text"},"type":"channels_select"},{"action_id":"VWAs","placeholder":{"emoji":true,"text":"Select the subject matter expert...","type":"plain_text"},"type":"users_select"}],"type":"actions"},{"block_id":"additional","element":{"action_id":"WKZt","multiline":true,"type":"plain_text_input"},"label":{"emoji":true,"text":"Provide any additional information:","type":"plain_text"},"optional":true,"type":"input"}],"bot_id":"B01B63T6ZFD","callback_id":"","clear_on_close":false,"close":{"emoji":true,"text":"Cancel","type":"plain_text"},"error":"","external_id":"","hash":"1602024115.StMLGPtB","id":"V01C3AAJMNW","notify_on_close":false,"ok":false,"previous_view_id":"","private_metadata":"incident","root_view_id":"V01C3AAJMNW","state":{"values":{"additional":{"WKZt":{"action_id":"","action_ts":"","block_id":"","initial_channel":"","initial_conversation":"","initial_date":"","initial_option":{"text":null,"value":""},"initial_user":"","selected_channel":"","selected_channels":null,"selected_conversation":"","selected_conversations":null,"selected_date":"","selected_option":{"text":null,"value":""},"selected_options":null,"selected_user":"","selected_users":null,"text":{"text":"","type":""},"type":"plain_text_input","value":"whoa"}},"bugzilla":{"VejTS":{"action_id":"","action_ts":"","block_id":"","initial_channel":"","initial_conversation":"","initial_date":"","initial_option":{"text":null,"value":""},"initial_user":"","selected_channel":"","selected_channels":null,"selected_conversation":"","selected_conversations":null,"selected_date":"","selected_option":{"text":null,"value":""},"selected_options":null,"selected_user":"","selected_users":null,"text":{"text":"","type":""},"type":"plain_text_input","value":"bug"}},"impact":{"fKxhJ":{"action_id":"","action_ts":"","block_id":"","initial_channel":"","initial_conversation":"","initial_date":"","initial_option":{"text":null,"value":""},"initial_user":"","selected_channel":"","selected_channels":null,"selected_conversation":"","selected_conversations":null,"selected_date":"","selected_option":{"text":null,"value":""},"selected_options":null,"selected_user":"","selected_users":null,"text":{"text":"","type":""},"type":"plain_text_input","value":"impact"}},"selectors":{"VWAs":{"action_id":"","action_ts":"","block_id":"","initial_channel":"","initial_conversation":"","initial_date":"","initial_option":{"text":null,"value":""},"initial_user":"","selected_channel":"","selected_channels":null,"selected_conversation":"","selected_conversations":null,"selected_date":"","selected_option":{"text":null,"value":""},"selected_options":null,"selected_user":"U01AN9288GP","selected_users":null,"text":{"text":"","type":""},"type":"users_select","value":""},"cJ/":{"action_id":"","action_ts":"","block_id":"","initial_channel":"","initial_conversation":"","initial_date":"","initial_option":{"text":null,"value":""},"initial_user":"","selected_channel":"C01B31AT7K4","selected_channels":null,"selected_conversation":"","selected_conversations":null,"selected_date":"","selected_option":{"text":null,"value":""},"selected_options":null,"selected_user":"","selected_users":null,"text":{"text":"","type":""},"type":"channels_select","value":""}},"summary":{"lEC":{"action_id":"","action_ts":"","block_id":"","initial_channel":"","initial_conversation":"","initial_date":"","initial_option":{"text":null,"value":""},"initial_user":"","selected_channel":"","selected_channels":null,"selected_conversation":"","selected_conversations":null,"selected_date":"","selected_option":{"text":null,"value":""},"selected_options":null,"selected_user":"","selected_users":null,"text":{"text":"","type":""},"type":"plain_text_input","value":"summary"}},"title":{"vyRE":{"action_id":"","action_ts":"","block_id":"","initial_channel":"","initial_conversation":"","initial_date":"","initial_option":{"text":null,"value":""},"initial_user":"","selected_channel":"","selected_channels":null,"selected_conversation":"","selected_conversations":null,"selected_date":"","selected_option":{"text":null,"value":""},"selected_options":null,"selected_user":"","selected_users":null,"text":{"text":"","type":""},"type":"plain_text_input","value":"title"}}}},"submit":{"emoji":true,"text":"Submit","type":"plain_text"},"team_id":"T01B37EA9T5","title":{"emoji":true,"text":"Document an Incident","type":"plain_text"},"type":"modal"}}`),
			blockIds: []string{"title", "summary", "impact", "additional", "bugzilla", "selectors"},
			expected: map[string]string{
				"title":                     "title",
				"summary":                   "summary",
				"impact":                    "impact",
				"additional":                "whoa",
				"bugzilla":                  "bug",
				"selectors_channels_select": "C01B31AT7K4",
				"selectors_users_select":    "U01AN9288GP",
			},
		},
		{
			name:     "complex input with a static selector",
			input:    []byte(`{"action_id":"","action_ts":"","actions":[],"api_app_id":"A01BJF00CAD","attachment_id":"","block_id":"","callback_id":"","channel":{"created":0,"creator":"","id":"","is_archived":false,"is_channel":false,"is_ext_shared":false,"is_general":false,"is_group":false,"is_im":false,"is_member":false,"is_mpim":false,"is_open":false,"is_org_shared":false,"is_pending_ext_shared":false,"is_private":false,"is_shared":false,"locale":"","members":null,"name":"","name_normalized":"","num_members":0,"priority":0,"purpose":{"creator":"","last_set":0,"value":""},"topic":{"creator":"","last_set":0,"value":""},"unlinked":0,"user":""},"container":{"attachment_id":0,"channel_id":"","is_app_unfurl":false,"is_ephemeral":false,"message_ts":"","type":"","view_id":""},"hash":"","is_cleared":false,"message":{"blocks":null,"delete_original":false,"replace_original":false},"message_ts":"","name":"","original_message":{"blocks":null,"delete_original":false,"replace_original":false},"response_url":"","submission":null,"team":{"domain":"dptp-robot-testing","id":"T01B37EA9T5","name":""},"token":"Rn1sVT099UjUeKclB3PkLT8t","trigger_id":"1445212076624.1377252349923.6db817174c902b85a7eac0d8f5613d3e","type":"view_submission","user":{"color":"","deleted":false,"enterprise_user":{"enterprise_id":"","enterprise_name":"","id":"","is_admin":false,"is_owner":false,"teams":null},"has_2fa":false,"has_files":false,"id":"U01B31ARZDG","is_admin":false,"is_app_user":false,"is_bot":false,"is_invited_user":false,"is_owner":false,"is_primary_owner":false,"is_restricted":false,"is_stranger":false,"is_ultra_restricted":false,"locale":"","name":"skuznets","presence":"","profile":{"display_name":"","display_name_normalized":"","email":"","fields":[],"first_name":"","image_192":"","image_24":"","image_32":"","image_48":"","image_72":"","image_original":"","last_name":"","phone":"","real_name":"","real_name_normalized":"","skype":"","status_expiration":0,"team":"","title":""},"real_name":"","team_id":"T01B37EA9T5","tz_label":"","tz_offset":0,"updated":0},"value":"","view":{"app_id":"A01BJF00CAD","blocks":[{"block_id":"kS/","text":{"emoji":true,"text":"Use this form to report a bug in the test platform or infrastructure.","type":"plain_text"},"type":"section"},{"accessory":{"action_id":"hqO4","text":{"emoji":true,"text":"Ask a Question","type":"plain_text"},"type":"button","value":"question"},"block_id":"nKpi","text":{"emoji":true,"text":"Please be certain that what you are reporting is a bug in the system. If it's not clear, please ask a question from the Test Platform Help-Desk engineer using the question form instead.","type":"plain_text"},"type":"section"},{"block_id":"2YZkQ","type":"divider"},{"block_id":"title","element":{"action_id":"EU7e8","type":"plain_text_input"},"label":{"emoji":true,"text":"Provide a title for this bug:","type":"plain_text"},"type":"input"},{"block_id":"category","element":{"action_id":"Ceww","options":[{"text":{"emoji":true,"text":"CI Jobs","type":"plain_text"},"value":"jobs"},{"text":{"emoji":true,"text":"CI Search","type":"plain_text"},"value":"search"},{"text":{"emoji":true,"text":"Release Controller","type":"plain_text"},"value":"Release Controller"},{"text":{"emoji":true,"text":"Other","type":"plain_text"},"value":"other"}],"placeholder":{"emoji":true,"text":"Select a category...","type":"plain_text"},"type":"static_select"},"label":{"emoji":true,"text":"What test infrastructure component is affected?","type":"plain_text"},"type":"input"},{"block_id":"optional","element":{"action_id":"DCgA","type":"plain_text_input"},"label":{"emoji":true,"text":"If other, what best describes the bugged component?","type":"plain_text"},"optional":true,"type":"input"},{"block_id":"yoR","type":"divider"},{"block_id":"symptom","element":{"action_id":"G=Vl","multiline":true,"type":"plain_text_input"},"label":{"emoji":true,"text":"What incorrect behavior did you notice?","type":"plain_text"},"type":"input"},{"block_id":"expected","element":{"action_id":"BFSg5","multiline":true,"type":"plain_text_input"},"label":{"emoji":true,"text":"What behavior did you expect instead?","type":"plain_text"},"type":"input"},{"block_id":"impact","element":{"action_id":"+2dWl","multiline":true,"type":"plain_text_input"},"label":{"emoji":true,"text":"What is the impact of this bug? How many jobs or users are impacted?","type":"plain_text"},"type":"input"},{"block_id":"reproduction","element":{"action_id":"Cv9","multiline":true,"type":"plain_text_input"},"label":{"emoji":true,"text":"Is this bug reproducible? If so, how?","type":"plain_text"},"type":"input"}],"bot_id":"B01B63T6ZFD","callback_id":"","clear_on_close":false,"close":{"emoji":true,"text":"Cancel","type":"plain_text"},"error":"","external_id":"","hash":"1602467948.iWjOP1Z9","id":"V01BYJ3JXN3","notify_on_close":false,"ok":false,"previous_view_id":"","private_metadata":"bug","root_view_id":"V01BYJ3JXN3","state":{"values":{"category":{"Ceww":{"action_id":"","action_ts":"","block_id":"","initial_channel":"","initial_conversation":"","initial_date":"","initial_option":{"text":null,"value":""},"initial_user":"","selected_channel":"","selected_channels":null,"selected_conversation":"","selected_conversations":null,"selected_date":"","selected_option":{"text":{"emoji":true,"text":"Other","type":"plain_text"},"value":"Other"},"selected_options":null,"selected_user":"","selected_users":null,"text":{"text":"","type":""},"type":"static_select","value":"Other"}},"expected":{"BFSg5":{"action_id":"","action_ts":"","block_id":"","initial_channel":"","initial_conversation":"","initial_date":"","initial_option":{"text":null,"value":""},"initial_user":"","selected_channel":"","selected_channels":null,"selected_conversation":"","selected_conversations":null,"selected_date":"","selected_option":{"text":null,"value":""},"selected_options":null,"selected_user":"","selected_users":null,"text":{"text":"","type":""},"type":"plain_text_input","value":"Something right!"}},"impact":{"+2dWl":{"action_id":"","action_ts":"","block_id":"","initial_channel":"","initial_conversation":"","initial_date":"","initial_option":{"text":null,"value":""},"initial_user":"","selected_channel":"","selected_channels":null,"selected_conversation":"","selected_conversations":null,"selected_date":"","selected_option":{"text":null,"value":""},"selected_options":null,"selected_user":"","selected_users":null,"text":{"text":"","type":""},"type":"plain_text_input","value":"I'm on fire."}},"optional":{"DCgA":{"action_id":"","action_ts":"","block_id":"","initial_channel":"","initial_conversation":"","initial_date":"","initial_option":{"text":null,"value":""},"initial_user":"","selected_channel":"","selected_channels":null,"selected_conversation":"","selected_conversations":null,"selected_date":"","selected_option":{"text":null,"value":""},"selected_options":null,"selected_user":"","selected_users":null,"text":{"text":"","type":""},"type":"plain_text_input","value":"My Component"}},"reproduction":{"Cv9":{"action_id":"","action_ts":"","block_id":"","initial_channel":"","initial_conversation":"","initial_date":"","initial_option":{"text":null,"value":""},"initial_user":"","selected_channel":"","selected_channels":null,"selected_conversation":"","selected_conversations":null,"selected_date":"","selected_option":{"text":null,"value":""},"selected_options":null,"selected_user":"","selected_users":null,"text":{"text":"","type":""},"type":"plain_text_input","value":"Every time, just push the button."}},"symptom":{"G=Vl":{"action_id":"","action_ts":"","block_id":"","initial_channel":"","initial_conversation":"","initial_date":"","initial_option":{"text":null,"value":""},"initial_user":"","selected_channel":"","selected_channels":null,"selected_conversation":"","selected_conversations":null,"selected_date":"","selected_option":{"text":null,"value":""},"selected_options":null,"selected_user":"","selected_users":null,"text":{"text":"","type":""},"type":"plain_text_input","value":"Something wrong!"}},"title":{"EU7e8":{"action_id":"","action_ts":"","block_id":"","initial_channel":"","initial_conversation":"","initial_date":"","initial_option":{"text":null,"value":""},"initial_user":"","selected_channel":"","selected_channels":null,"selected_conversation":"","selected_conversations":null,"selected_date":"","selected_option":{"text":null,"value":""},"selected_options":null,"selected_user":"","selected_users":null,"text":{"text":"","type":""},"type":"plain_text_input","value":"My Title"}}}},"submit":{"emoji":true,"text":"Submit","type":"plain_text"},"team_id":"T01B37EA9T5","title":{"emoji":true,"text":"File a Bug","type":"plain_text"},"type":"modal"}}`),
			blockIds: []string{"question", "category", "optional", "symptom", "expected", "impact", "reproduction"},
			expected: map[string]string{
				"expected":               "Something right!",
				"impact":                 "I'm on fire.",
				"category_static_select": "Other",
				"optional":               "My Component",
				"reproduction":           "Every time, just push the button.",
				"symptom":                "Something wrong!",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var callback *slack.InteractionCallback
			if err := json.Unmarshal(testCase.input, &callback); err != nil {
				t.Errorf("%s: failed to unmarshal callback: %v", testCase.name, err)
				return
			}
			if diff := cmp.Diff(valuesFor(callback, testCase.blockIds...), testCase.expected); diff != "" {
				t.Errorf("%s: got incorrect values: %v", testCase.name, diff)
			}
		})
	}
}

func TestBulletListFunc(t *testing.T) {
	var testCases = []struct {
		input, output string
	}{
		{
			input:  "one line",
			output: "* one line",
		},
		{
			input:  "one line trailing newline\n",
			output: "* one line trailing newline",
		},
		{
			input:  "many\nlines",
			output: "* many\n* lines",
		},
		{
			input:  "many\n\n\n\n\n\n\n\nnewlines",
			output: "* many\n* newlines",
		},
	}
	toBulletList := BulletListFunc()["toBulletList"].(func(string) string)

	for _, testCase := range testCases {
		if diff := cmp.Diff(toBulletList(testCase.input), testCase.output); diff != "" {
			t.Errorf("got incorrect output: %v", diff)
		}
	}
}