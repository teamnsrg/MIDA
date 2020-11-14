package yibrowse

// A single argument (or return value) from an API call
type Arg struct {
	T   string `json:"type" bson:"type"`
	Val string `json:"val" bson:"val"`
}

// A single API call
type Call struct {
	T        string `json:"t" bson:"t"`
	C        string `json:"c" bson:"c"`
	F        string `json:"f" bson:"f"`
	Args     []Arg  `json:"a" bson:"a"`
	Ret      Arg    `json:"r" bson:"r"`
	ScriptId string `json:"-" bson:"-"`
}

// A single script, identified by a unique script ID
type Script struct {
	ScriptId string `json:"script_id" bson:"script_id"`
	Url      string `json:"base_url" bson:"base_url"`
	Calls    []Call `json:"calls" bson:"-"`
	SHA1     string `json:"sha1" bson:"-"`
	Length   int    `json:"length" bson:"-"`

	// Fingerprinting
	OpenWPM OpenWPMResults `json:"openwpm_results,omitempty" bson:"openwpm_results,omitempty"`
}

// The trace from a single isolate. Script IDs are only
// guaranteed unique per-isolate
type Isolate struct {
	Scripts map[string]*Script `json:"scripts" bson:"-"`

	// MongoDB-use only fields
	ID       int64   `json:"-" bson:"_id"`
	Parent   int64   `json:"-" bson:"parent"`
	Children []int64 `json:"-" bson:"children"`
}

// A full trace, parsed and ready to be stored or processed further
type RawJSTrace struct {
	Isolates map[string]*Isolate `json:"isolates,omitempty" bson:"-"`

	// Scripts for which we saw calls but never saw an initial declaration
	// We store this for use in repairing the trace using script metadata
	// UnknownScripts[isolate][scriptId] = true
	UnknownScripts map[string]map[string]bool `json:"-" bson:"-"`

	// Parsing data
	IgnoredCalls int `json:"ignored_calls"`
	StoredCalls  int `json:"stored_calls"`
}

type CleanedJSTrace struct {
	Url     string             `json:"url,omitempty"`
	Scripts map[string]*Script `json:"scripts,omitempty"`
}

type OpenWPMResults struct {
	Canvas     bool `json:"canvas,omitempty" bson:"canvas,omitempty"`
	CanvasFont bool `json:"canvas_font,omitempty" bson:"canvas_font,omitempty"`
	WebRTC     bool `json:"web_rtc,omitempty" bson:"web_rtc,omitempty"`
	Audio      bool `json:"audio,omitempty" bson:"audio,omitempty"`
	Battery    bool `json:"battery,omitempty" bson:"battery,omitempty"`
}

type LineType int

const (
	ErrorLine   LineType = -1
	UnknownLine LineType = 0
	CallLine    LineType = 1
	ArgLine     LineType = 2
	RetLine     LineType = 3
	OtherLine   LineType = 4
	ControlLine LineType = 5
)

type Line struct {
	LT         LineType
	LineNum    int
	Isolate    string
	IsRet      bool
	ArgType    string
	ArgVal     string
	CallType   string
	CallClass  string
	CallFunc   string
	BaseURL    string
	ScriptId   string
	IsCallback bool
	IsBegin    bool
	TS         string // Timestamp
}
