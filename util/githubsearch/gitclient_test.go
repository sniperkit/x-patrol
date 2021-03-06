/*

Copyright (c) 2018 sec.xiaomi.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THEq
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

*/

package githubsearch_test

import (
	"testing"
	"../githubsearch"
	"fmt"
)

//func TestClient_GetUserInfo(t *testing.T) {
//	t.Log(githubsearch.GithubClient.GetUserInfo("netxfly"))
//}
//
//func TestClient_GetOrgsMembers(t *testing.T) {
//	t.Log(githubsearch.GithubClient.GetOrgsMembers("netxfly"))
//}
//
//func TestClient_GetOrgsRepos(t *testing.T) {
//	t.Log(githubsearch.GithubClient.GetOrgsRepos("xsec-lab"))
//}
//
//func TestClient_GetUserRepos(t *testing.T) {
//	t.Log(githubsearch.GithubClient.GetUserRepos("netxfly"))
//}
//
//func TestClient_GetUserOrgs(t *testing.T) {
//	t.Log(githubsearch.GithubClient.GetUserOrgs("54chen"))
//}

func TestSearchCode(t *testing.T) {
	gitClient, _, _ := githubsearch.GetGithubClient()
	codeSearchResults, _ := gitClient.SearchCode("spdb")
	for _, codeSearchResult := range codeSearchResults {
		for _, codeResult := range codeSearchResult.CodeResults {
			fmt.Println(codeResult.TextMatches)
			fmt.Println(codeResult.HTMLURL)
		}
	}
}

func TestBuildQuery(t *testing.T) {
	query := "shang"
	buildedQuery, err := githubsearch.BuildQuery(query)
	if err == nil {
		fmt.Println(buildedQuery)
	}
}