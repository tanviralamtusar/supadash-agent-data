# SupaDash Dashboard — Task Checklist

## Phase 1: Studio Fork & Patch System ✅
- [x] Create `studio/` directory structure
- [x] Create [studio/build.sh](file:///d:/Coding/supadash/SupaDash/studio/build.sh) — Download + patch + build
- [x] Create [studio/patch.sh](file:///d:/Coding/supadash/SupaDash/studio/patch.sh) — Apply patch files
- [x] Create [studio/Dockerfile](file:///d:/Coding/supadash/SupaDash/studio/Dockerfile) — Build custom Studio image
- [x] Create patch: [01-api-urls.patch](file:///d:/Coding/supadash/SupaDash/studio/patches/01-api-urls.patch) — Redirect API calls
- [x] Create patch: [02-auth-integration.patch](file:///d:/Coding/supadash/SupaDash/studio/patches/02-auth-integration.patch) — Use SupaDash auth
- [x] Create patch: [03-remove-cloud-features.patch](file:///d:/Coding/supadash/SupaDash/studio/patches/03-remove-cloud-features.patch) — Strip cloud UI  
- [x] Create patch: [04-branding.patch](file:///d:/Coding/supadash/SupaDash/studio/patches/04-branding.patch) — SupaDash logos + title
- [x] Add `supadash-studio` to [docker-compose.yaml](file:///d:/Coding/supadash/SupaDash/docker-compose.yaml)
- [x] Copy branding assets to `files/public/img/`

## Phase 2: Auth Integration + 2FA
- [x] Patch Studio login to call SupaDash `/auth/token` (in Phase 1 patches)
- [x] Add TOTP 2FA endpoints to Go API
- [x] Add database migration for 2FA secrets
- [x] Add 2FA setup page in Studio
- [x] Test: login → 2FA → dashboard

## Phase 3: Core Dashboard Pages
- [x] Patch: Project List page
- [x] Patch: Create Project page (with plan selector)
- [x] Patch: Project Dashboard
- [x] Patch: Table Editor / SQL Editor
- [x] Patch: Auth Users, Storage, Edge Functions, Logs

## Phase 4: Resource Manager ✅
- [x] Create Resource Manager page (`/project/[ref]/resources`)
- [x] Implement UI components (CPU/RAM Gauges, Service Table, Scaling Sliders)
- [x] Implement Recommendations and Burst Pool UI
- [x] Integrate Admin Server Overview page (`/server/resources`)
- [x] Connect scaling & plan update controls to Go backend
- [x] Patch Studio sidebar to include "Resources" link
- [x] Verify API integration and real-time updates

## Phase 5: Team Management ✅
- [x] Patch team page for SupaDash API
- [x] Implement organization members query hook
- [x] Implement organization invitation mutation hook
- [x] Implement organization member role update/delete hooks
- [x] Update backend `GetOrganizationMembers` query with `gotrue_id`


## Phase 6: Real-time Updates ✅
- [x] Add `gorilla/websocket` to Go backend
- [x] Implement `WsHub` and `/ws` endpoint in Go API
- [x] Broadcast project status and container stats via WebSocket
- [x] Implement `useRealtime` hook in Studio frontend
- [x] Patch Studio to use real-time updates for project state

## Phase 7: Branding & Polish ✅
- [x] Update Global Metadata (`_app.tsx`)
- [x] Update Navigation Branding (`LayoutHeader.tsx`, `HomeIcon.tsx`)
- [x] Update Auth Page Branding (`SignInLayout.tsx`, `SignInPartner.tsx`, `SignUpForm.tsx`)
- [x] Update Support & Help Branding (`HelpOptionsList.tsx`, `SupportFormPage.tsx`, `AIAssistantOption.tsx`, `DiscordCTACard.tsx`)
- [x] Update User Dropdown Branding (`UserDropdown.tsx`)
- [x] Audit & Brand Core Interface Components
    - [x] Storage Settings (`StorageSettings.tsx`)
    - [x] API Settings (`PostgrestConfig.tsx`, `DataApiProjectUrlCard.tsx`)
    - [x] Database Settings (`SettingsDatabaseEmptyStateLocal.tsx`)
    - [x] Authentication Settings (`AuthProvidersForm/index.tsx`)
    - [x] Table Editor Components (`GridHeaderActions.tsx`, `SidePanelEditor.tsx`)
- [x] Final UI Cleanup
    - [x] Brand "Getting Started" section (`GettingStarted.tsx`, `GettingStarted.utils.tsx`)
    - [x] Global "Supabase" text removal (Final pass)
- [x] Verification & QA
    - [x] Verify real-time functionality
    - [x] Verify help/support redirects
