// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/middlewares"
	"github.com/rsvix/go-htmx-app-template/internal/structs"
	"strconv"
)

func SnippetsPage(c echo.Context, pageTitle string, userName string, snippets []structs.Snippet) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"bg-white dark:bg-gray-900 antialiased m-2 md:m-4 lg:m-6 xl:m-8 2xl:m-10\"><!-- Page title --><div class=\"text-center p-0 mt-0 mb-2 md:mb-4 lg:mb-6 xl:mb-8 2xl:mb-10\"><h2 class=\"text-3xl m-0 pb-2 font-bold leading-tight tracking-tight text-gray-900 dark:text-white\">Snippets page</h2></div><!-- Page content --><div class=\"flow-root\"><div class=\"bg-white dark:bg-gray-800 relative shadow-md sm:rounded-lg overflow-hidden\"><!-- Table header --><div class=\"flex flex-col md:flex-row items-center justify-between space-y-3 md:space-y-0 md:space-x-4 p-4\"><div class=\"w-full md:w-1/2\"><form class=\"flex items-center\"><label for=\"simple-search\" class=\"sr-only\">Search</label><div class=\"relative w-full\"><div class=\"absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none\"><svg aria-hidden=\"true\" class=\"w-5 h-5 text-gray-500 dark:text-gray-400\" fill=\"currentColor\" viewbox=\"0 0 20 20\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" d=\"M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z\" clip-rule=\"evenodd\"></path></svg></div><input type=\"text\" id=\"simple-search\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full pl-10 p-2 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500\" placeholder=\"Search\" required=\"\"></div></form></div><div class=\"w-full md:w-auto flex flex-col md:flex-row space-y-2 md:space-y-0 items-stretch md:items-center justify-end md:space-x-3 flex-shrink-0\"><button type=\"button\" class=\"flex items-center justify-center text-white bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-primary-600 dark:hover:bg-primary-700 focus:outline-none dark:focus:ring-primary-800\"><svg class=\"h-3.5 w-3.5 mr-2\" fill=\"currentColor\" viewbox=\"0 0 20 20\" xmlns=\"http://www.w3.org/2000/svg\" aria-hidden=\"true\"><path clip-rule=\"evenodd\" fill-rule=\"evenodd\" d=\"M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z\"></path></svg> Add snippet</button><div class=\"flex items-center space-x-3 w-full md:w-auto\"><button id=\"actionsDropdownButton\" data-dropdown-toggle=\"actionsDropdown\" class=\"w-full md:w-auto flex items-center justify-center py-2 px-4 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-primary-700 focus:z-10 focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700\" type=\"button\"><svg class=\"-ml-1 mr-1.5 w-5 h-5\" fill=\"currentColor\" viewbox=\"0 0 20 20\" xmlns=\"http://www.w3.org/2000/svg\" aria-hidden=\"true\"><path clip-rule=\"evenodd\" fill-rule=\"evenodd\" d=\"M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z\"></path></svg> Actions</button><div id=\"actionsDropdown\" class=\"hidden z-10 w-44 bg-white rounded divide-y divide-gray-100 shadow dark:bg-gray-700 dark:divide-gray-600\"><ul class=\"py-1 text-sm text-gray-700 dark:text-gray-200\" aria-labelledby=\"actionsDropdownButton\"><li><a href=\"#\" class=\"block py-2 px-4 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Mass Edit</a></li></ul><div class=\"py-1\"><a href=\"#\" class=\"block py-2 px-4 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white\">Delete all</a></div></div><button id=\"filterDropdownButton\" data-dropdown-toggle=\"filterDropdown\" class=\"w-full md:w-auto flex items-center justify-center py-2 px-4 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-primary-700 focus:z-10 focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700\" type=\"button\"><svg xmlns=\"http://www.w3.org/2000/svg\" aria-hidden=\"true\" class=\"h-4 w-4 mr-2 text-gray-400\" viewbox=\"0 0 20 20\" fill=\"currentColor\"><path fill-rule=\"evenodd\" d=\"M3 3a1 1 0 011-1h12a1 1 0 011 1v3a1 1 0 01-.293.707L12 11.414V15a1 1 0 01-.293.707l-2 2A1 1 0 018 17v-5.586L3.293 6.707A1 1 0 013 6V3z\" clip-rule=\"evenodd\"></path></svg> Filter <svg class=\"-mr-1 ml-1.5 w-5 h-5\" fill=\"currentColor\" viewbox=\"0 0 20 20\" xmlns=\"http://www.w3.org/2000/svg\" aria-hidden=\"true\"><path clip-rule=\"evenodd\" fill-rule=\"evenodd\" d=\"M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z\"></path></svg></button><div id=\"filterDropdown\" class=\"z-10 hidden w-48 p-3 bg-white rounded-lg shadow dark:bg-gray-700\"><h6 class=\"mb-3 text-sm font-medium text-gray-900 dark:text-white\">Choose brand</h6><ul class=\"space-y-2 text-sm\" aria-labelledby=\"filterDropdownButton\"><li class=\"flex items-center\"><input id=\"apple\" type=\"checkbox\" value=\"\" class=\"w-4 h-4 bg-gray-100 border-gray-300 rounded text-primary-600 focus:ring-primary-500 dark:focus:ring-primary-600 dark:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500\"> <label for=\"apple\" class=\"ml-2 text-sm font-medium text-gray-900 dark:text-gray-100\">Apple (56)</label></li><li class=\"flex items-center\"><input id=\"fitbit\" type=\"checkbox\" value=\"\" class=\"w-4 h-4 bg-gray-100 border-gray-300 rounded text-primary-600 focus:ring-primary-500 dark:focus:ring-primary-600 dark:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500\"> <label for=\"fitbit\" class=\"ml-2 text-sm font-medium text-gray-900 dark:text-gray-100\">Microsoft (16)</label></li><li class=\"flex items-center\"><input id=\"razor\" type=\"checkbox\" value=\"\" class=\"w-4 h-4 bg-gray-100 border-gray-300 rounded text-primary-600 focus:ring-primary-500 dark:focus:ring-primary-600 dark:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500\"> <label for=\"razor\" class=\"ml-2 text-sm font-medium text-gray-900 dark:text-gray-100\">Razor (49)</label></li><li class=\"flex items-center\"><input id=\"nikon\" type=\"checkbox\" value=\"\" class=\"w-4 h-4 bg-gray-100 border-gray-300 rounded text-primary-600 focus:ring-primary-500 dark:focus:ring-primary-600 dark:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500\"> <label for=\"nikon\" class=\"ml-2 text-sm font-medium text-gray-900 dark:text-gray-100\">Nikon (12)</label></li><li class=\"flex items-center\"><input id=\"benq\" type=\"checkbox\" value=\"\" class=\"w-4 h-4 bg-gray-100 border-gray-300 rounded text-primary-600 focus:ring-primary-500 dark:focus:ring-primary-600 dark:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500\"> <label for=\"benq\" class=\"ml-2 text-sm font-medium text-gray-900 dark:text-gray-100\">BenQ (74)</label></li></ul></div></div></div></div><!-- Table --><div class=\"overflow-x-auto\"><table class=\"w-full text-sm text-left text-gray-500 dark:text-gray-400\"><!-- Table head --><thead class=\"text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400\"><tr><th scope=\"col\" class=\"px-4 py-3\">Snippet name</th><th scope=\"col\" class=\"px-4 py-3\">Category</th><th scope=\"col\" class=\"px-4 py-3\">Language</th><th scope=\"col\" class=\"px-4 py-3\">Owner</th><th scope=\"col\" class=\"px-4 py-3\">View</th><th scope=\"col\" class=\"px-4 py-3\">Edit</th></tr></thead><!-- Table body --><tbody><!-- Loop -->")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, snippet := range snippets {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr class=\"border-b dark:border-gray-700\"><th scope=\"row\" class=\"px-4 py-3 font-medium text-gray-900 whitespace-nowrap dark:text-white\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(snippet.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/snippets.templ`, Line: 117, Col: 148}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</th><td class=\"px-4 py-3\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(snippet.Category)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/snippets.templ`, Line: 118, Col: 80}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"px-4 py-3\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(snippet.Language)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/snippets.templ`, Line: 119, Col: 80}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"px-4 py-3\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(snippet.Owner)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/snippets.templ`, Line: 120, Col: 77}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><!-- View --><td class=\"px-4 py-3\"><button hx-get=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs("/snippetview/" + strconv.FormatUint(uint64(snippet.Id), 10))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/snippets.templ`, Line: 124, Col: 115}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"body\" hx-swap=\"beforeend\" class=\"inline-block rounded hover:bg-indigo-700 cursor-pointer text-xl\"><i class=\"fa-regular fa-eye\"></i></button></td><!-- Edit --><td class=\"px-4 py-3\"><button hx-get=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs("/snippetedit/" + strconv.FormatUint(uint64(snippet.Id), 10))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/snippets.templ`, Line: 135, Col: 115}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"body\" hx-swap=\"beforeend\" class=\"inline-block rounded hover:bg-indigo-700 cursor-pointer text-xl\"><i class=\"fa-regular fa-pen-to-square\"></i></button></td></tr>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tbody></table></div><!-- Table footer --><nav class=\"flex flex-col md:flex-row justify-between items-start md:items-center space-y-3 md:space-y-0 p-4\" aria-label=\"Table navigation\"><span class=\"text-sm font-normal text-gray-500 dark:text-gray-400\">Showing <span class=\"font-semibold text-gray-900 dark:text-white\">1-10</span> of <span class=\"font-semibold text-gray-900 dark:text-white\">1000</span></span><ul class=\"inline-flex items-stretch -space-x-px\"><li><a href=\"#\" class=\"flex items-center justify-center h-full py-1.5 px-3 ml-0 text-gray-500 bg-white rounded-l-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white\"><span class=\"sr-only\">Previous</span> <svg class=\"w-5 h-5\" aria-hidden=\"true\" fill=\"currentColor\" viewbox=\"0 0 20 20\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" d=\"M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z\" clip-rule=\"evenodd\"></path></svg></a></li><li><a href=\"#\" class=\"flex items-center justify-center text-sm py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white\">1</a></li><li><a href=\"#\" class=\"flex items-center justify-center text-sm py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white\">2</a></li><li><a href=\"#\" aria-current=\"page\" class=\"flex items-center justify-center text-sm z-10 py-2 px-3 leading-tight text-primary-600 bg-primary-50 border border-primary-300 hover:bg-primary-100 hover:text-primary-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white\">3</a></li><li><a href=\"#\" class=\"flex items-center justify-center text-sm py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white\">...</a></li><li><a href=\"#\" class=\"flex items-center justify-center text-sm py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white\">100</a></li><li><a href=\"#\" class=\"flex items-center justify-center h-full py-1.5 px-3 leading-tight text-gray-500 bg-white rounded-r-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white\"><span class=\"sr-only\">Next</span> <svg class=\"w-5 h-5\" aria-hidden=\"true\" fill=\"currentColor\" viewbox=\"0 0 20 20\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" d=\"M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z\" clip-rule=\"evenodd\"></path></svg></a></li></ul></nav></div></div></section><script type=\"text/javascript\" nonce=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(middlewares.GetRandomNonce(c))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/snippets.templ`, Line: 193, Col: 76}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">\n            document.addEventListener('htmx:afterRequest', function(evt) {\n                hljs.highlightAll();\n            });\n        </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Layout(c, pageTitle, true, userName, true).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
