package main

/*
Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

Example 1:

Input: s = "egg", t = "add"

Output: true

Explanation:

The strings s and t can be made identical by:

Mapping 'e' to 'a'.
Mapping 'g' to 'd'.
Example 2:

Input: s = "foo", t = "bar"

Output: false

Explanation:

The strings s and t can not be made identical as 'o' needs to be mapped to both 'a' and 'r'.

Example 3:

Input: s = "paper", t = "title"

Output: true

Constraints:

1 <= s.length <= 5 * 104
t.length == s.length
s and t consist of any valid ascii character.
*/

// isIsomorphic関数: 文字列 s と t を受け取り、bool型 (true/false) を返す
func isIsomorphic(s string, t string) bool {
	// m: sの文字からtの文字へのマッピングを保持するマップ (TypeScriptの Map<string, string> や { [key: string]: string } に相当)
	// make(map[byte]byte) でキーも値もbyte型(1バイト文字)のマップを作成
	m := make(map[byte]byte)
	// mapped: tの文字が既にマッピング先として使われたかどうかを記録するマップ (TypeScriptの Set<string> に相当)
	// 値に struct{} (空の構造体) を使うのは、キーの存在だけを記録したい場合のGoの慣用的な方法です。メモリを消費しません。
	mapped := make(map[byte]struct{})

	// s の各文字についてループ (TypeScriptの for (let i = 0; i < s.length; i++) に相当)
	for i := 0; i < len(s); i++ {
		// s[i]: sのi番目の文字 (byte型)
		// m[s[i]] でマップmからキー s[i] に対応する値を取得しようと試みる
		// Goでは、マップのアクセスは2つの値を返す: 1つ目は値 (存在しない場合はゼロ値)、2つ目はキーが存在したかどうか(bool)
		// c に値、ok に存在有無(true/false)が入る (TypeScriptの Map.has() と Map.get() を組み合わせたような挙動)
		if c, ok := m[s[i]]; ok {
			// マップ m に s[i] のマッピングが既に存在する場合 (ok == true)
			// 既存のマッピング c が、現在の t[i] と一致するかチェック
			if c != t[i] {
				// 一致しない場合、ルール1違反なので false を返す
				return false
			}
		} else {
			// マップ m に s[i] のマッピングが存在しない場合 (ok == false)
			// 次に、t[i] が既に他の文字のマッピング先として使われていないかチェック
			// mapped[t[i]] でマップmappedにキー t[i] が存在するかチェック
			// _ (ブランク識別子) は、値は不要で存在有無(ok)だけを知りたい場合に使用
			if _, ok := mapped[t[i]]; ok {
				// 既に使われている場合 (ok == true)、ルール2違反なので false を返す
				return false
			}
			// s[i] から t[i] への新しいマッピングをマップmに記録 (TypeScriptの m.set(s[i], t[i]) や m[s[i]] = t[i] に相当)
			m[s[i]] = t[i]
			// t[i] を「マッピング先として使用済み」としてマップmappedに記録 (TypeScriptの mapped.add(t[i]) に相当)
			// struct{}{} は空の構造体のインスタンスを作成して値として入れている
			mapped[t[i]] = struct{}{}
		}
	}
	// ループが最後まで完了した場合、矛盾がなかったので true を返す
	return true
}
