package util

import (
	"github.com/itzngga/goRoxy/helper"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func CreateHydratedTemplateButton(content, footer string, buttons ...*waProto.HydratedTemplateButton) *waProto.Message {
	return &waProto.Message{
		TemplateMessage: &waProto.TemplateMessage{
			HydratedTemplate: &waProto.HydratedFourRowTemplate{
				HydratedContentText: proto.String(content),
				HydratedFooterText:  proto.String(footer),
				HydratedButtons:     buttons,
			},
		},
	}
}

func GenerateHydratedUrlButton(text, url string) *waProto.HydratedTemplateButton {
	return &waProto.HydratedTemplateButton{
		HydratedButton: &waProto.HydratedTemplateButton_UrlButton{
			UrlButton: &waProto.HydratedURLButton{
				DisplayText: proto.String(text),
				Url:         proto.String(url),
			},
		},
	}
}

func GenerateHydratedCallButton(text, number string) *waProto.HydratedTemplateButton {
	return &waProto.HydratedTemplateButton{
		HydratedButton: &waProto.HydratedTemplateButton_CallButton{
			CallButton: &waProto.HydratedCallButton{
				DisplayText: proto.String(text),
				PhoneNumber: proto.String(number),
			},
		},
	}
}

func GenerateHydratedQuickReplyButton(text, id, cmd string) *waProto.HydratedTemplateButton {
	return &waProto.HydratedTemplateButton{
		HydratedButton: &waProto.HydratedTemplateButton_QuickReplyButton{
			QuickReplyButton: &waProto.HydratedQuickReplyButton{
				DisplayText: proto.String(text),
				Id:          proto.String(helper.CreateButtonID(id, cmd)),
			},
		},
	}
}
