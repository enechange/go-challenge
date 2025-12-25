# REPORT

## 技術的決定

### 1. アーキテクチャ: 3層構造

**決定**: handler / service / repository の3層構造を採用

```
handler/    - HTTPリクエスト/レスポンス処理
service/    - ビジネスロジック（距離計算、フィルタリング）
repository/ - データアクセス（sqlc生成コードのラッパー）
```

**理由**:
- 各層の責務が明確に分離される
- テスタビリティが高い（各層をモック可能）
- 課題要件の「テスタビリティを意識した設計」に適合

**トレードオフ**:
- シンプルな課題に対してはやや過剰な構造かもしれない
- ファイル数が増える

---

### 2. 開発手法: TDD (Test-Driven Development)

**決定**: t-wadaさん方式のTDD（Red-Green-Refactor）で実装

**理由**:
- テストファーストにより、テスト可能な設計が自然と生まれる
- 小さなステップで確実に進められる
- リグレッションを防げる

---

## 問題点・課題

（実装進行中に追記予定）

---

## 参考

- [Go-Challenge Interface Specification Document](./Go-Challenge%20Interface%20Specification%20Document.pdf)
