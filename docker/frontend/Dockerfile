FROM node:latest

WORKDIR /app/frontend

# アプリケーションの依存関係ファイルをコピー
COPY /app/frontend/package.json ./
COPY /app/frontend/package-lock.json ./

# 依存関係をインストール
RUN npm install

# アプリケーションのビルド（本番環境用の場合）
# RUN npm run build

ENTRYPOINT ["npm", "start"]