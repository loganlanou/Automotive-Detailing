# Deployment Guide

This guide explains how to deploy Detailing Pass to [Vercel](https://vercel.com) using the included `vercel.json` configuration and the serverless handler in `api/index.go`.

## Why Vercel?

- ✅ Free tier with automatic HTTPS and global CDN
- ✅ Git-powered deployments with preview environments
- ✅ Native Go Serverless Functions (used by `api/index.go`)
- ✅ Automatic Tailwind/Templ build via `build.sh`
- ✅ Instant rollbacks and shareable preview URLs

## Step-by-Step Deployment

### 1. Push to GitHub

```bash
git add -A
git commit -m "Prepare for deployment"
git push origin main
```

Vercel will pull straight from GitHub, so make sure `main` (or your chosen branch) is up-to-date.

### 2. Create a Vercel Project

1. Sign in at [vercel.com](https://vercel.com) with GitHub
2. Click **Add New… → Project**
3. Import the `Automotive-Detailing` repository
4. Leave the root directory as `/`
5. When prompted for the framework, choose **Other**

The project will automatically pick up the checked-in `vercel.json`, so you do not need to edit build settings in the UI.

### 3. Build & Install Commands

`vercel.json` already specifies:

- `installCommand`: `npm install` (installs Tailwind/PostCSS)
- `buildCommand`: `bash build.sh`
- `framework`: `null` (disables opinionated defaults)

`build.sh` performs the following:

1. Installs `templ`
2. Generates Templ output (`templ generate`)
3. Compiles Tailwind to `web/static/css/output.css`
4. Builds the binary (useful for local/docker deploys)

The Vercel deployment itself relies on `api/index.go`, so the built binary is optional but harmless.

### 4. Environment Variables

For the serverless build, the in-memory SQLite database defined in `api/index.go` is used, so no variables are required. If you later wire the project to an external database or need API keys (Clerk, SMTP, etc.), add them under **Settings → Environment Variables** in the Vercel dashboard and redeploy.

### 5. Domains & URLs

1. After the first deploy, click **Visit** to open the preview URL
2. Promote to production with **Deploy to Production**
3. Add a custom domain under **Settings → Domains** if desired

### 6. Logs & Monitoring

- Build logs live under the **Deployments** tab
- Runtime logs for `api/index.go` are visible via **Functions → Logs**
- `/health` route returns `{ "status": "ok" }` for uptime checks

### 7. Troubleshooting

**Build fails**
- Ensure Node 18+ features are not required (Vercel uses Node 18 by default)
- Confirm `templ` is reachable (Vercel builders have Go preinstalled)
- Run `bash build.sh` locally to catch syntax or Tailwind errors

**Runtime errors**
- Check the **Functions** tab for stack traces
- Verify that `web/static/css/output.css` exists in git (it is generated during build and should be committed if you prefer deterministic builds)

**Styling missing**
- Make sure Tailwind input/output paths in `build.sh` line up with your CSS entrypoint

## Alternative Deployment Options

Prefer a traditional VM/container workflow? Use one of these instead:

### Render.com
1. Sign up at [render.com](https://render.com)
2. Create a new Web Service
3. Build command: `make build`
4. Start command: `./bin/server`

### Fly.io
1. Install flyctl: `curl -L https://fly.io/install.sh | sh`
2. Run `flyctl launch`
3. Configure the generated `fly.toml` and deploy

Both options keep SQLite on persistent storage if you need write access, whereas Vercel’s serverless handler resets state between invocations.
