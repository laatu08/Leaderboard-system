FROM node:18-alpine

WORKDIR /app

COPY package.json package-lock.json* ./
RUN npm install

COPY . .

# Build Expo Web
RUN npx expo export --platform web

# Install a simple static server
RUN npm install -g serve

EXPOSE 3000

CMD ["serve", "-s", "dist", "-l", "3000"]
