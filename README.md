# Artifex AI: Full-Stack AI-Embedded Social Networking Platform

Artifex AI is a scalable, full-stack social networking application built with **React** (frontend), **Go microservices** (backend), and deployed on **Google App Engine**. The platform integrates the **DALL·E 3 API** for AI-generated images and features real-time search, personalized feeds, and performant caching — making it ideal for exploring modern web architecture and AI-enhanced social tooling.

---

## 🚀 Features

### 🔧 Core Functionality
- **User Feed & Interactions:** Users can create posts, like, comment, follow, and explore feeds.
- **AI-Generated Content:** DALL·E 3 integration powers image creation for posts.
- **Realtime Search:** Elasticsearch cluster supports full-text search with custom analyzers & filters.
- **Authentication:** JWT-based auth with refresh token rotation.
- **Caching:** Redis caching for posts & user profiles improves performance.

### 📈 Performance Highlights
- Sub-second full-text search across **20,000+ posts**.
- Backend serving **400+ concurrent users** with **p95 latency < 200 ms**.
- Feed rendering time reduced from ~1.2 s to ~400 ms via caching.

---

## 🛠 Tech Stack

| Layer | Technology |
|------|------------|
| Frontend | React, Ant Design |
| Backend | Go, RESTful APIs |
| Search | Elasticsearch |
| Deployment | Google App Engine |
| Caching | Redis |
| AI & Images | OpenAI DALL·E 3 API |

---

## 📦 Getting Started

### Prerequisites

Make sure you have the following installed:

- Node.js (v16+)
- Go (v1.20+)
- Docker (optional — for local services)
- Google Cloud SDK

### Clone & Install

```bash
git clone https://github.com/yunbo2614/ArtifexAI.git
cd SocialAI
