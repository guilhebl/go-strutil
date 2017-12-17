package strutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Tests filtering a text with HTML
func TestFilterHtmlTags(t *testing.T) {
	s := `
<p>Links:</p><ul><li><a href="foo">Foo</a><li>
<a href="/bar/baz">BarBaz</a></ul><span>TEXT <b>I</b> WANT</span>
<script type='text/javascript'>
/* <![CDATA[ */
var post_notif_widget_ajax_obj = {"ajax_url":"http:\/\/site.com\/wp-admin\/admin-ajax.php","nonce":"9b8270e2ef","processing_msg":"Processing..."};
/* ]]> */
</script>`

	text1 := FilterHtmlTags(s)
	assert.Equal(t, "Links:FooBarBazTEXTIWANT", text1)

	s2 := `
	<!DOCTYPE html PUBLIC "-//IETF//DTD HTML 2.0//EN">
	<HTML>
	   <HEAD>
		  <TITLE>
			 A Small Hello
		  </TITLE>
	   </HEAD>
	<BODY>
	<H1>Hi</H1><P>This is very minimal "hello world" HTML document.</P><div><h3>content</h3></div>
	</BODY>
	</HTML>`
	text2 := FilterHtmlTags(s2)

	assert.Equal(t, "A Small HelloHiThis is very minimal \"hello world\" HTML document.content", text2)
}

// test filtering a text with no HTML tags
func TestFilterHtmlNoTags(t *testing.T) {
	s := "This is a simple test 1234567890 with no tags!"
	text := FilterHtmlTags(s)
	assert.Equal(t, s, text)
}

// test trimming a last char suffix
func TestTrimSuffix(t *testing.T) {
	s := "tree+abc+def+gef+"
	text := TrimSuffix(s, "+")
	assert.Equal(t, "tree+abc+def+gef", text)
}

// test trimming a last char empty suffix expect no change
func TestTrimSuffixEmptySuffix(t *testing.T) {
	s := "tree+abc+def+gef+"
	text := TrimSuffix(s, "")
	assert.Equal(t, s, text)
}

// test trimming a last char sequence suffix
func TestTrimSuffixCharSequence(t *testing.T) {
	s := "tree+abc+def+gef+"
	text := TrimSuffix(s, "+def+gef+")
	assert.Equal(t, "tree+abc", text)
}

// test Combining String array with separator
func TestCombineStrings(t *testing.T) {
	strings := []string{"orange", "apple", "banana", "grape"}
	s := CombineStrings(strings, "-")
	assert.Equal(t, "orange-apple-banana-grape", s)
}

// test MakeString using array with blank separator
func TestMakeString(t *testing.T) {
	strings := []string{"orange", "apple", "banana", "grape"}
	s := MakeString(strings)
	assert.Equal(t, "orangeapplebananagrape", s)
}
