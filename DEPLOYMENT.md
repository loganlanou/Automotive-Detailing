# Deployment Guide

This guide will help you deploy Detailing Pass to Railway.app for free.

## Why Railway?

Railway is perfect for Go applications:
- ✅ Free tier available (500 hours/month)
- ✅ Automatic builds from Git
- ✅ Instant public URL
- ✅ Built-in database support
- ✅ Zero configuration needed

## Step-by-Step Deployment

### 1. Push to GitHub

First, make sure your code is pushed to GitHub:

```bash
git add -A
git commit -m "Prepare for deployment"
git push origin main
```

### 2. Sign Up for Railway

1. Go to [railway.app](https://railway.app)
2. Click "Login" in the top right
3. Sign up with your GitHub account
4. Authorize Railway to access your repositories

### 3. Deploy Your Project

1. Click "New Project" on your Railway dashboard
2. Select "Deploy from GitHub repo"
3. Choose your `Automotive-Detailing` repository
4. Railway will automatically detect it's a Go project and start building

### 4. Configure Environment Variables (Optional)

Railway will use sensible defaults, but you can customize:

1. Click on your deployed service
2. Go to "Variables" tab
3. Add any custom variables:
   - `PORT` - Auto-set by Railway
   - `DATABASE_PATH` - Defaults to `./data/detailing.db`

### 5. Get Your Public URL

1. Go to "Settings" tab
2. Click "Generate Domain" under "Networking"
3. Your site will be available at: `https://your-project-name.up.railway.app`

### 6. Share with Your Friend!

Copy the URL and send it to anyone. The site will be:
- ✅ Live 24/7
- ✅ Publicly accessible
- ✅ Using the free tier
- ✅ Auto-deploying on git push

## What's Included

The deployment includes:
- ✅ Automatic database creation
- ✅ All static assets (CSS, JS, images)
- ✅ Ford F-150 inventory images from Courtesy Auto
- ✅ Responsive design that works on mobile

## Troubleshooting

### Build Fails
- Check the build logs in Railway dashboard
- Ensure all dependencies are in `go.mod`
- Verify `nixpacks.toml` is in project root

### Site Not Loading
- Check if service is "Active" in Railway dashboard
- Verify domain is generated in Settings
- Check deployment logs for errors

### Database Issues
- Railway automatically creates the database on first run
- Schema is embedded in the binary
- Data persists across deployments

## Alternative Deployment Options

If you prefer other platforms:

### Render.com
1. Sign up at [render.com](https://render.com)
2. Create new "Web Service"
3. Connect your GitHub repo
4. Use build command: `make build`
5. Use start command: `./bin/server`

### Fly.io
1. Install flyctl: `curl -L https://fly.io/install.sh | sh`
2. Sign up: `flyctl auth signup`
3. Deploy: `flyctl launch`
4. Follow the prompts

## Cost

Railway free tier includes:
- 500 hours per month (enough for 24/7 operation)
- $5 credit per month
- After free tier: ~$5/month for this app

## Support

If you run into issues:
- Check Railway docs: [docs.railway.app](https://docs.railway.app)
- Railway Discord: [discord.gg/railway](https://discord.gg/railway)
