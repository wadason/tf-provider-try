# tf-provider-try

## このページは何？？

terraform providerの実装方法を学ぶ。
まだ存在しないSaaSのproviderを実装することが目標。

## 参考

- https://www.hashicorp.com/blog/writing-custom-terraform-providers

## build方法

```
go build -o terraform-provider-example
terraform plan
```

## TODO

- [ ] 簡単なapiを叩いてみる


## 実装メモ

- resourceが宣言できる状態にした
- resourceの変更を検知するためのsetIdを追加した
- resourceの更新失敗時に、ハンドリングができる状態にした

## やってみたい

- 下記を一通りやる
- https://www.terraform.io/docs/extend/how-terraform-works.html
- 適当なapiで実践
