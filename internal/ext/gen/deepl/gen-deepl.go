// Package deepl provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package deepl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

const (
	Auth_headerScopes = "auth_header.Scopes"
)

// Defines values for Formality.
const (
	Default    Formality = "default"
	Less       Formality = "less"
	More       Formality = "more"
	PreferLess Formality = "prefer_less"
	PreferMore Formality = "prefer_more"
)

// Defines values for SourceLanguage.
const (
	SourceLanguageBG SourceLanguage = "BG"
	SourceLanguageCS SourceLanguage = "CS"
	SourceLanguageDA SourceLanguage = "DA"
	SourceLanguageDE SourceLanguage = "DE"
	SourceLanguageEL SourceLanguage = "EL"
	SourceLanguageEN SourceLanguage = "EN"
	SourceLanguageES SourceLanguage = "ES"
	SourceLanguageET SourceLanguage = "ET"
	SourceLanguageFI SourceLanguage = "FI"
	SourceLanguageFR SourceLanguage = "FR"
	SourceLanguageHU SourceLanguage = "HU"
	SourceLanguageID SourceLanguage = "ID"
	SourceLanguageIT SourceLanguage = "IT"
	SourceLanguageJA SourceLanguage = "JA"
	SourceLanguageKO SourceLanguage = "KO"
	SourceLanguageLT SourceLanguage = "LT"
	SourceLanguageLV SourceLanguage = "LV"
	SourceLanguageNB SourceLanguage = "NB"
	SourceLanguageNL SourceLanguage = "NL"
	SourceLanguagePL SourceLanguage = "PL"
	SourceLanguagePT SourceLanguage = "PT"
	SourceLanguageRO SourceLanguage = "RO"
	SourceLanguageRU SourceLanguage = "RU"
	SourceLanguageSK SourceLanguage = "SK"
	SourceLanguageSL SourceLanguage = "SL"
	SourceLanguageSV SourceLanguage = "SV"
	SourceLanguageTR SourceLanguage = "TR"
	SourceLanguageUK SourceLanguage = "UK"
	SourceLanguageZH SourceLanguage = "ZH"
)

// Defines values for TargetLanguage.
const (
	TargetLanguageBG   TargetLanguage = "BG"
	TargetLanguageCS   TargetLanguage = "CS"
	TargetLanguageDA   TargetLanguage = "DA"
	TargetLanguageDE   TargetLanguage = "DE"
	TargetLanguageEL   TargetLanguage = "EL"
	TargetLanguageENGB TargetLanguage = "EN-GB"
	TargetLanguageENUS TargetLanguage = "EN-US"
	TargetLanguageES   TargetLanguage = "ES"
	TargetLanguageET   TargetLanguage = "ET"
	TargetLanguageFI   TargetLanguage = "FI"
	TargetLanguageFR   TargetLanguage = "FR"
	TargetLanguageHU   TargetLanguage = "HU"
	TargetLanguageID   TargetLanguage = "ID"
	TargetLanguageIT   TargetLanguage = "IT"
	TargetLanguageJA   TargetLanguage = "JA"
	TargetLanguageKO   TargetLanguage = "KO"
	TargetLanguageLT   TargetLanguage = "LT"
	TargetLanguageLV   TargetLanguage = "LV"
	TargetLanguageNB   TargetLanguage = "NB"
	TargetLanguageNL   TargetLanguage = "NL"
	TargetLanguagePL   TargetLanguage = "PL"
	TargetLanguagePTBR TargetLanguage = "PT-BR"
	TargetLanguagePTPT TargetLanguage = "PT-PT"
	TargetLanguageRO   TargetLanguage = "RO"
	TargetLanguageRU   TargetLanguage = "RU"
	TargetLanguageSK   TargetLanguage = "SK"
	TargetLanguageSL   TargetLanguage = "SL"
	TargetLanguageSV   TargetLanguage = "SV"
	TargetLanguageTR   TargetLanguage = "TR"
	TargetLanguageUK   TargetLanguage = "UK"
	TargetLanguageZH   TargetLanguage = "ZH"
)

// Defines values for GetLanguagesParamsType.
const (
	Source GetLanguagesParamsType = "source"
	Target GetLanguagesParamsType = "target"
)

// Defines values for TranslateTextFormdataBodyOutlineDetection.
const (
	TranslateTextFormdataBodyOutlineDetectionN0 TranslateTextFormdataBodyOutlineDetection = "0"
)

// Defines values for TranslateTextFormdataBodyPreserveFormatting.
const (
	TranslateTextFormdataBodyPreserveFormattingN0 TranslateTextFormdataBodyPreserveFormatting = "0"
	TranslateTextFormdataBodyPreserveFormattingN1 TranslateTextFormdataBodyPreserveFormatting = "1"
)

// Defines values for TranslateTextFormdataBodySplitSentences.
const (
	N0         TranslateTextFormdataBodySplitSentences = "0"
	N1         TranslateTextFormdataBodySplitSentences = "1"
	Nonewlines TranslateTextFormdataBodySplitSentences = "nonewlines"
)

// Defines values for TranslateTextFormdataBodyTagHandling.
const (
	Html TranslateTextFormdataBodyTagHandling = "html"
	Xml  TranslateTextFormdataBodyTagHandling = "xml"
)

// Formality Sets whether the translated text should lean towards formal or informal language.
// This feature currently only works for target languages
// `DE` (German),
// `FR` (French),
// `IT` (Italian),
// `ES` (Spanish),
// `NL` (Dutch),
// `PL` (Polish),
// `PT-PT`,
// `PT-BR` (Portuguese)
// and `RU` (Russian).
// Setting this parameter with a target language that does not support formality will fail,
// unless one of the `prefer_...` options are used.
// Possible options are:
//   - `default` (default)
//   - `more` - for a more formal language
//   - `less` - for a more informal language
//   - `prefer_more` - for a more formal language if available, otherwise fallback to default formality
//   - `prefer_less` - for a more informal language if available, otherwise fallback to default formality
type Formality string

// SourceLanguage Language of the text to be translated. Options currently available:
//   - `BG` - Bulgarian
//   - `CS` - Czech
//   - `DA` - Danish
//   - `DE` - German
//   - `EL` - Greek
//   - `EN` - English
//   - `ES` - Spanish
//   - `ET` - Estonian
//   - `FI` - Finnish
//   - `FR` - French
//   - `HU` - Hungarian
//   - `ID` - Indonesian
//   - `IT` - Italian
//   - `JA` - Japanese
//   - `KO` - Korean
//   - `LT` - Lithuanian
//   - `LV` - Latvian
//   - `NB` - Norwegian (Bokmål)
//   - `NL` - Dutch
//   - `PL` - Polish
//   - `PT` - Portuguese (all Portuguese varieties mixed)
//   - `RO` - Romanian
//   - `RU` - Russian
//   - `SK` - Slovak
//   - `SL` - Slovenian
//   - `SV` - Swedish
//   - `TR` - Turkish
//   - `UK` - Ukrainian
//   - `ZH` - Chinese
//
// If this parameter is omitted, the API will attempt to detect the language of the text and translate it.
type SourceLanguage string

// TargetLanguage The language into which the text should be translated. Options currently available:
//   - `BG` - Bulgarian
//   - `CS` - Czech
//   - `DA` - Danish
//   - `DE` - German
//   - `EL` - Greek
//   - `EN` - English (unspecified variant for backward compatibility; please select `EN-GB` or `EN-US` instead)
//   - `EN-GB` - English (British)
//   - `EN-US` - English (American)
//   - `ES` - Spanish
//   - `ET` - Estonian
//   - `FI` - Finnish
//   - `FR` - French
//   - `HU` - Hungarian
//   - `ID` - Indonesian
//   - `IT` - Italian
//   - `JA` - Japanese
//   - `KO` - Korean
//   - `LT` - Lithuanian
//   - `LV` - Latvian
//   - `NB` - Norwegian (Bokmål)
//   - `NL` - Dutch
//   - `PL` - Polish
//   - `PT` - Portuguese (unspecified variant for backward compatibility; please select `PT-BR` or `PT-PT` instead)
//   - `PT-BR` - Portuguese (Brazilian)
//   - `PT-PT` - Portuguese (all Portuguese varieties excluding Brazilian Portuguese)
//   - `RO` - Romanian
//   - `RU` - Russian
//   - `SK` - Slovak
//   - `SL` - Slovenian
//   - `SV` - Swedish
//   - `TR` - Turkish
//   - `UK` - Ukrainian
//   - `ZH` - Chinese (simplified)
type TargetLanguage string

// GetLanguagesParams defines parameters for GetLanguages.
type GetLanguagesParams struct {
	// Type Sets whether source or target languages should be listed. Possible options are:
	//  * `source` (default): For languages that can be used in the `source_lang` parameter of [translate](https://www.deepl.com/docs-api/translate-text/translate-text) requests.
	//  * `target`: For languages that can be used in the `target_lang` parameter of [translate](https://www.deepl.com/docs-api/translate-text/translate-text) requests.
	Type *GetLanguagesParamsType `form:"type,omitempty" json:"type,omitempty"`
}

// GetLanguagesParamsType defines parameters for GetLanguages.
type GetLanguagesParamsType string

// TranslateTextFormdataBody defines parameters for TranslateText.
type TranslateTextFormdataBody struct {
	// Formality Sets whether the translated text should lean towards formal or informal language.
	// This feature currently only works for target languages
	// `DE` (German),
	// `FR` (French),
	// `IT` (Italian),
	// `ES` (Spanish),
	// `NL` (Dutch),
	// `PL` (Polish),
	// `PT-PT`,
	// `PT-BR` (Portuguese)
	// and `RU` (Russian).
	// Setting this parameter with a target language that does not support formality will fail,
	// unless one of the `prefer_...` options are used.
	// Possible options are:
	//   * `default` (default)
	//   * `more` - for a more formal language
	//   * `less` - for a more informal language
	//   * `prefer_more` - for a more formal language if available, otherwise fallback to default formality
	//   * `prefer_less` - for a more informal language if available, otherwise fallback to default formality
	Formality *Formality `json:"formality,omitempty"`

	// GlossaryId Specify the glossary to use for the translation. **Important:** This requires the `source_lang` parameter to be set and the language pair of the glossary has to match the language pair of the request.
	GlossaryId *string `json:"glossary_id,omitempty"`

	// IgnoreTags Comma-separated list of XML tags that indicate text not to be translated.
	//
	//
	// Use this paramter to ensure that elements in the original text are not altered in the translation (e.g., trademarks, product names) and insert tags into your original text. In the following example, the `ignore_tags` parameter is set to `keep`:
	//  * Example request:
	//    ```
	//    Please open the page <keep>Settings</keep> to configure your system.
	//    ```
	//  * Example response:
	//    ```
	//    Bitte öffnen Sie die Seite <keep>Settings</keep> um Ihr System zu konfigurieren.
	//    ```
	IgnoreTags *string `json:"ignore_tags,omitempty"`

	// NonSplittingTags Comma-separated list of XML tags which never split sentences.
	//
	//
	// For some XML files, finding tags with textual content and splitting sentences using those tags won't yield the best results. The following example shows the engine splitting sentences on `par` tags and proceeding to translate the parts separately, resulting in an incorrect translation:
	//  * Example request:
	// ```
	// <par>The firm said it had been </par><par> conducting an internal investigation.</par>
	// ```
	//  * Example response:
	// ```
	// <par>Die Firma sagte, es sei eine gute Idee gewesen.</par><par> Durchführung einer internen Untersuchung.</par>
	// ```
	//
	//
	// As this can lead to bad translations, this type of structure should either be avoided, or the `non_splitting_tags` parameter should be set. The following example shows the same call, with the parameter set to `par`:
	//  * Example request:
	// ```
	// <par>The firm said it had been </par><par> conducting an internal investigation.</par>
	// ```
	//  * Example response:
	// ```
	// <par>Die Firma sagte, dass sie</par><par> eine interne Untersuchung durchgeführt</par><par> habe</par><par>.</par>
	// ```
	//
	//
	// This time, the sentence is translated as a whole. The XML tags are now considered markup and copied into the translated sentence. As the translation of the words "had been" has moved to another position in the German sentence, the two par tags are duplicated (which is expected here).
	NonSplittingTags *string `json:"non_splitting_tags,omitempty"`

	// OutlineDetection The automatic detection of the XML structure won't yield best results in all XML files. You can disable this automatic mechanism altogether by setting the `outline_detection` parameter to `0` and selecting the tags that should be considered structure tags. This will split sentences using the `splitting_tags` parameter.
	//
	//
	// In the example below, we achieve the same results as the automatic engine by disabling automatic detection with `outline_detection=0` and setting the parameters manually to `tag_handling=xml`, `split_sentences=nonewlines`,  and `splitting_tags=par,title`.
	//  * Example request:
	//    ```
	//    <document>
	//      <meta>
	//        <title>A document's title</title>
	//      </meta>
	//      <content>
	//        <par>This is the first sentence. Followed by a second one.</par>
	//        <par>This is the third sentence.</par>
	//      </content>
	//    </document>
	//    ```
	//  * Example response:
	//    ```
	//    <document>
	//      <meta>
	//        <title>Der Titel eines Dokuments</title>
	//      </meta>
	//      <content>
	//        <par>Das ist der erste Satz. Gefolgt von einem zweiten.</par>
	//        <par>Dies ist der dritte Satz.</par>
	//      </content>
	//    </document>
	//    ```
	// While this approach is slightly more complicated, it allows for greater control over the structure of the translation output.
	OutlineDetection *TranslateTextFormdataBodyOutlineDetection `json:"outline_detection,omitempty"`

	// PreserveFormatting Sets whether the translation engine should respect the original formatting, even if it would usually correct some aspects. Possible values are:
	//  * `0` (default)
	//  * `1`
	//
	// The formatting aspects affected by this setting include:
	//  * Punctuation at the beginning and end of the sentence
	//  * Upper/lower case at the beginning of the sentence
	PreserveFormatting *TranslateTextFormdataBodyPreserveFormatting `json:"preserve_formatting,omitempty"`

	// SourceLang Language of the text to be translated. Options currently available:
	//  * `BG` - Bulgarian
	//  * `CS` - Czech
	//  * `DA` - Danish
	//  * `DE` - German
	//  * `EL` - Greek
	//  * `EN` - English
	//  * `ES` - Spanish
	//  * `ET` - Estonian
	//  * `FI` - Finnish
	//  * `FR` - French
	//  * `HU` - Hungarian
	//  * `ID` - Indonesian
	//  * `IT` - Italian
	//  * `JA` - Japanese
	//  * `KO` - Korean
	//  * `LT` - Lithuanian
	//  * `LV` - Latvian
	//  * `NB` - Norwegian (Bokmål)
	//  * `NL` - Dutch
	//  * `PL` - Polish
	//  * `PT` - Portuguese (all Portuguese varieties mixed)
	//  * `RO` - Romanian
	//  * `RU` - Russian
	//  * `SK` - Slovak
	//  * `SL` - Slovenian
	//  * `SV` - Swedish
	//  * `TR` - Turkish
	//  * `UK` - Ukrainian
	//  * `ZH` - Chinese
	//
	// If this parameter is omitted, the API will attempt to detect the language of the text and translate it.
	SourceLang *SourceLanguage `json:"source_lang,omitempty"`

	// SplitSentences Sets whether the translation engine should first split the input into sentences. For text translations where
	// `tag_handling` is not set to `html`, the default value is `1`, meaning the engine splits on punctuation and on newlines.
	//
	// For text translations where `tag_handling=html`, the default value is `nonewlines`, meaning the engine splits on punctuation only, ignoring newlines.
	//
	// The use of `nonewlines` as the default value for text translations where `tag_handling=html` is new behavior that was implemented in November 2022,
	// when HTML handling was moved out of beta.
	//
	// Possible values are:
	//
	//  * `0` - no splitting at all, whole input is treated as one sentence
	//  * `1` (default when `tag_handling` is not set to `html`) - splits on punctuation and on newlines
	//  * `nonewlines` (default when `tag_handling=html`) - splits on punctuation only, ignoring newlines
	//
	// For applications that send one sentence per text parameter, we recommend setting `split_sentences` to `0`, in order to prevent the engine from splitting the sentence unintentionally.
	//
	//
	// Please note that newlines will split sentences when `split_sentences=1`. We recommend cleaning files so they don't contain breaking sentences or setting the parameter `split_sentences` to `nonewlines`.
	SplitSentences *TranslateTextFormdataBodySplitSentences `json:"split_sentences,omitempty"`

	// SplittingTags Comma-separated list of XML tags which always cause splits.
	//
	//
	// See the example in the `outline_detection` parameter's description.
	SplittingTags *string `json:"splitting_tags,omitempty"`

	// TagHandling Sets which kind of tags should be handled. Options currently available:
	//  * `xml`: Enable XML tag handling; see [XML Handling](https://www.deepl.com/docs-api/xml).
	//  * `html`: Enable HTML tag handling; see [HTML Handling](https://www.deepl.com/docs-api/html).
	TagHandling *TranslateTextFormdataBodyTagHandling `json:"tag_handling,omitempty"`

	// TargetLang The language into which the text should be translated. Options currently available:
	//  * `BG` - Bulgarian
	//  * `CS` - Czech
	//  * `DA` - Danish
	//  * `DE` - German
	//  * `EL` - Greek
	//  * `EN` - English (unspecified variant for backward compatibility; please select `EN-GB` or `EN-US` instead)
	//  * `EN-GB` - English (British)
	//  * `EN-US` - English (American)
	//  * `ES` - Spanish
	//  * `ET` - Estonian
	//  * `FI` - Finnish
	//  * `FR` - French
	//  * `HU` - Hungarian
	//  * `ID` - Indonesian
	//  * `IT` - Italian
	//  * `JA` - Japanese
	//  * `KO` - Korean
	//  * `LT` - Lithuanian
	//  * `LV` - Latvian
	//  * `NB` - Norwegian (Bokmål)
	//  * `NL` - Dutch
	//  * `PL` - Polish
	//  * `PT` - Portuguese (unspecified variant for backward compatibility; please select `PT-BR` or `PT-PT` instead)
	//  * `PT-BR` - Portuguese (Brazilian)
	//  * `PT-PT` - Portuguese (all Portuguese varieties excluding Brazilian Portuguese)
	//  * `RO` - Romanian
	//  * `RU` - Russian
	//  * `SK` - Slovak
	//  * `SL` - Slovenian
	//  * `SV` - Swedish
	//  * `TR` - Turkish
	//  * `UK` - Ukrainian
	//  * `ZH` - Chinese (simplified)
	TargetLang TargetLanguage `json:"target_lang"`

	// Text Text to be translated. Only UTF-8-encoded plain text is supported. The parameter may be specified multiple times and translations are returned in the same order as they are requested. Each of the parameter values may contain multiple sentences.
	Text string `json:"text"`
}

// TranslateTextFormdataBodyOutlineDetection defines parameters for TranslateText.
type TranslateTextFormdataBodyOutlineDetection string

// TranslateTextFormdataBodyPreserveFormatting defines parameters for TranslateText.
type TranslateTextFormdataBodyPreserveFormatting string

// TranslateTextFormdataBodySplitSentences defines parameters for TranslateText.
type TranslateTextFormdataBodySplitSentences string

// TranslateTextFormdataBodyTagHandling defines parameters for TranslateText.
type TranslateTextFormdataBodyTagHandling string

// TranslateTextFormdataRequestBody defines body for TranslateText for application/x-www-form-urlencoded ContentType.
type TranslateTextFormdataRequestBody TranslateTextFormdataBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetLanguages request
	GetLanguages(ctx context.Context, params *GetLanguagesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TranslateText request with any body
	TranslateTextWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	TranslateTextWithFormdataBody(ctx context.Context, body TranslateTextFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetUsage request
	GetUsage(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetLanguages(ctx context.Context, params *GetLanguagesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLanguagesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TranslateTextWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTranslateTextRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TranslateTextWithFormdataBody(ctx context.Context, body TranslateTextFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTranslateTextRequestWithFormdataBody(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetUsage(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUsageRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetLanguagesRequest generates requests for GetLanguages
func NewGetLanguagesRequest(server string, params *GetLanguagesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/languages")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Type != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "type", runtime.ParamLocationQuery, *params.Type); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewTranslateTextRequestWithFormdataBody calls the generic TranslateText builder with application/x-www-form-urlencoded body
func NewTranslateTextRequestWithFormdataBody(server string, body TranslateTextFormdataRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	bodyStr, err := runtime.MarshalForm(body, nil)
	if err != nil {
		return nil, err
	}
	bodyReader = strings.NewReader(bodyStr.Encode())
	return NewTranslateTextRequestWithBody(server, "application/x-www-form-urlencoded", bodyReader)
}

// NewTranslateTextRequestWithBody generates requests for TranslateText with any type of body
func NewTranslateTextRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/translate")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetUsageRequest generates requests for GetUsage
func NewGetUsageRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/usage")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetLanguages request
	GetLanguagesWithResponse(ctx context.Context, params *GetLanguagesParams, reqEditors ...RequestEditorFn) (*GetLanguagesResponse, error)

	// TranslateText request with any body
	TranslateTextWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*TranslateTextResponse, error)

	TranslateTextWithFormdataBodyWithResponse(ctx context.Context, body TranslateTextFormdataRequestBody, reqEditors ...RequestEditorFn) (*TranslateTextResponse, error)

	// GetUsage request
	GetUsageWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUsageResponse, error)
}

type GetLanguagesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		// Language The language code of the given language.
		Language string `json:"language"`

		// Name Name of the language in English.
		Name string `json:"name"`

		// SupportsFormality Denotes formality support in case of a target language listing.
		SupportsFormality *bool `json:"supports_formality,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetLanguagesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLanguagesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type TranslateTextResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Translations *[]struct {
			// DetectedSourceLanguage The language detected in the source text. It reflects the value of the `source_lang` parameter, when specified.
			DetectedSourceLanguage *SourceLanguage `json:"detected_source_language,omitempty"`

			// Text The translated text.
			Text *string `json:"text,omitempty"`
		} `json:"translations,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r TranslateTextResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TranslateTextResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUsageResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// CharacterCount Characters translated so far in the current billing period.
		CharacterCount *int64 `json:"character_count,omitempty"`

		// CharacterLimit Current maximum number of characters that can be translated per billing period.
		CharacterLimit *int64 `json:"character_limit,omitempty"`

		// DocumentCount Current maximum number of documents that can be translated per billing period.
		DocumentCount *int64 `json:"document_count,omitempty"`

		// DocumentLimit Documents translated so far in the current billing period.
		DocumentLimit *int64 `json:"document_limit,omitempty"`

		// TeamDocumentCount Current maximum number of documents that can be translated by the team per billing period.
		TeamDocumentCount *int64 `json:"team_document_count,omitempty"`

		// TeamDocumentLimit Documents translated by all users in the team so far in the current billing period.
		TeamDocumentLimit *int64 `json:"team_document_limit,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetUsageResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUsageResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetLanguagesWithResponse request returning *GetLanguagesResponse
func (c *ClientWithResponses) GetLanguagesWithResponse(ctx context.Context, params *GetLanguagesParams, reqEditors ...RequestEditorFn) (*GetLanguagesResponse, error) {
	rsp, err := c.GetLanguages(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLanguagesResponse(rsp)
}

// TranslateTextWithBodyWithResponse request with arbitrary body returning *TranslateTextResponse
func (c *ClientWithResponses) TranslateTextWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*TranslateTextResponse, error) {
	rsp, err := c.TranslateTextWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseTranslateTextResponse(rsp)
}

func (c *ClientWithResponses) TranslateTextWithFormdataBodyWithResponse(ctx context.Context, body TranslateTextFormdataRequestBody, reqEditors ...RequestEditorFn) (*TranslateTextResponse, error) {
	rsp, err := c.TranslateTextWithFormdataBody(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseTranslateTextResponse(rsp)
}

// GetUsageWithResponse request returning *GetUsageResponse
func (c *ClientWithResponses) GetUsageWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUsageResponse, error) {
	rsp, err := c.GetUsage(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUsageResponse(rsp)
}

// ParseGetLanguagesResponse parses an HTTP response from a GetLanguagesWithResponse call
func ParseGetLanguagesResponse(rsp *http.Response) (*GetLanguagesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLanguagesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			// Language The language code of the given language.
			Language string `json:"language"`

			// Name Name of the language in English.
			Name string `json:"name"`

			// SupportsFormality Denotes formality support in case of a target language listing.
			SupportsFormality *bool `json:"supports_formality,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseTranslateTextResponse parses an HTTP response from a TranslateTextWithResponse call
func ParseTranslateTextResponse(rsp *http.Response) (*TranslateTextResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TranslateTextResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Translations *[]struct {
				// DetectedSourceLanguage The language detected in the source text. It reflects the value of the `source_lang` parameter, when specified.
				DetectedSourceLanguage *SourceLanguage `json:"detected_source_language,omitempty"`

				// Text The translated text.
				Text *string `json:"text,omitempty"`
			} `json:"translations,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetUsageResponse parses an HTTP response from a GetUsageWithResponse call
func ParseGetUsageResponse(rsp *http.Response) (*GetUsageResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUsageResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// CharacterCount Characters translated so far in the current billing period.
			CharacterCount *int64 `json:"character_count,omitempty"`

			// CharacterLimit Current maximum number of characters that can be translated per billing period.
			CharacterLimit *int64 `json:"character_limit,omitempty"`

			// DocumentCount Current maximum number of documents that can be translated per billing period.
			DocumentCount *int64 `json:"document_count,omitempty"`

			// DocumentLimit Documents translated so far in the current billing period.
			DocumentLimit *int64 `json:"document_limit,omitempty"`

			// TeamDocumentCount Current maximum number of documents that can be translated by the team per billing period.
			TeamDocumentCount *int64 `json:"team_document_count,omitempty"`

			// TeamDocumentLimit Documents translated by all users in the team so far in the current billing period.
			TeamDocumentLimit *int64 `json:"team_document_limit,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}