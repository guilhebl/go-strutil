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
