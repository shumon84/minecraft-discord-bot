# death-reporter
Minecraft のログを監視し、ユーザが死亡した場合に死亡ログを Discord の特定のテキストチャンネルに送信する bot

#　Getting Started
``` shell
# カレントディレクトリに "config.json" という名前で設定ファイルを事前に用意しておく。 
# 詳細な設定項目については pkg/infrastructure/config/... を確認すること。
export CONFIG_PATH="./config.json" # config ファイルの path を設定する

go install github.com/shumon84/minecraft-discord-bot/cmd/death-reporter
death-reporter
```
