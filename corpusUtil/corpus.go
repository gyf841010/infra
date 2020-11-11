package corpusUtil

import (
	"encoding/json"
	. "github.com/gyf841010/infra/logging"
	"github.com/gyf841010/infra/commonUtil"
	"github.com/gyf841010/infra/baseUtil"
	"github.com/gyf841010/infra/httpUtil"
)

const (
	POST = "POST"
	GET  = "GET"

	FACE_MODULE     = "face"
	CLOTHES_MODULE  = "clothes"
	ACTION_MODULE   = "action"
	BIND_MODULE     = "bind"
	SMART_FURNITURE = "smart_furniture"
)

var DEFAULT_CORPUS = baseUtil.CorpusData{
	Start: baseUtil.SingleCorpus{
		Text:   "这个嘛，暂时我还不会，等我会了再表演给你吧",
		Action: "",
	},
}

// Get Corpus Request
type CorpusRequest struct {
	Module   string `json:"module" description:"功能模块"`
	CaseName string `json:"caseName" description:"查询条件"`
}

type CorpusList struct {
	Text   []string `json:"text" description:"回复语料文本"`
	Action []string `json:"action" description:"回复语料动作"`
}

// Corpus
type CorpusData struct {
	Id          int        `json:"id" description:"唯一id"`
	ApiId       int        `json:"apiId" description:"api id"`
	CaseName    string     `json:"caseName" description:"查询条件"`
	Description string     `json:"description" description:"描叙"`
	Start       CorpusList `json:"startCorpus" description:"开始语料"`
	Exec        CorpusList `json:"execCorpus" description:"执行语料"`
	End         CorpusList `json:"endCorpus" description:"结束语料"`
}

// Corpus Response
type CorpusResponse struct {
	ResCode int          `json:"code" required:"true" description:"0成功 >1失败"`
	Message string       `json:"message" description:"提示消息"`
	Data    []CorpusData `json:"data" description:"语料数据"`
}

func GetCorpus(corpusReq *CorpusRequest, url string) *baseUtil.CorpusData {
	if corpusReq.CaseName == "" {
		return &DEFAULT_CORPUS
	}
	var corpus baseUtil.CorpusData

	respContent, err := httpUtil.PostJson(url, nil, corpusReq)
	if err != nil {
		Log.Error("Failed to Post Corpus Sever ", WithError(err))
		return &DEFAULT_CORPUS
	}

	if len(respContent) <= 0 {
		Log.Error("Failed To Read Resp Body while len is 0", WithError(err))
		return &DEFAULT_CORPUS
	}
	Log.Debug("Corpus Response", With("bodyBytes", string(respContent)))

	corpusResponse := &CorpusResponse{}
	if err = json.Unmarshal(respContent, corpusResponse); err != nil {
		Log.Error("Failed To Read Resp Body while Unmarshal err", With("body", string(respContent)), WithError(err))
		return &DEFAULT_CORPUS
	}
	Log.Debug("Corpus Response", With("corpusResponse", corpusResponse))

	if len(corpusResponse.Data) <= 0 {
		return &DEFAULT_CORPUS
	}

	if len(corpusResponse.Data[0].Start.Text) > 0 {
		randomID := commonUtil.GetRandom(len(corpusResponse.Data[0].Start.Text))
		corpus.Start.Text = corpusResponse.Data[0].Start.Text[randomID]
	}
	if len(corpusResponse.Data[0].Start.Action) > 0 {
		randomID := commonUtil.GetRandom(len(corpusResponse.Data[0].Start.Action))
		corpus.Start.Action = corpusResponse.Data[0].Start.Action[randomID]
	}

	if len(corpusResponse.Data[0].Exec.Text) > 0 {
		randomID := commonUtil.GetRandom(len(corpusResponse.Data[0].Exec.Text))
		corpus.Exec.Text = corpusResponse.Data[0].Exec.Text[randomID]
	}
	if len(corpusResponse.Data[0].Exec.Action) > 0 {
		randomID := commonUtil.GetRandom(len(corpusResponse.Data[0].Exec.Action))
		corpus.Exec.Action = corpusResponse.Data[0].Exec.Action[randomID]
	}

	if len(corpusResponse.Data[0].End.Text) > 0 {
		randomID := commonUtil.GetRandom(len(corpusResponse.Data[0].End.Text))
		corpus.End.Text = corpusResponse.Data[0].End.Text[randomID]
	}
	if len(corpusResponse.Data[0].End.Action) > 0 {
		randomID := commonUtil.GetRandom(len(corpusResponse.Data[0].End.Action))
		corpus.End.Action = corpusResponse.Data[0].End.Action[randomID]
	}

	return &corpus

}
