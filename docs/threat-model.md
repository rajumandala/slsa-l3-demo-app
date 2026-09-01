# Threat Model — stages, attack classes, real incidents

This table is the spine of the hands-on labs in this repo. Each row maps a
supply-chain stage to an attack class and a real, named incident, and each
one is addressed by a specific lab and control built in this repo — see the
"Addressed by" column, filled in as each module lands.

| Stage | Attack class | Real example | Addressed by |
|---|---|---|---|
| Source | Trusted-maintainer social engineering, backdoored commits | xz-utils backdoor, CVE-2024-3094 (2024) | Module 1 — Source Track L3/L4 |
| Source | Maintainer account takeover via phishing | npm `debug`/`chalk` maintainer phishing (Sept 2025) | Module 1 — Source Track L2/L3 (branch protection, signed commits) |
| Build | Build system compromise, code injected at compile time then signed legitimately | SolarWinds Orion (2020) | Module 2 — Build Track L3 |
| Build | CI misconfiguration -> cache poisoning -> malicious artifact published under legitimate identity | Ultralytics PyPI compromise via GitHub Actions cache poisoning (Dec 2024) | Module 2 — Build Track L2 isolation + hardening baseline |
| Build | Compromised bot credential -> mutable tag rewrite -> secrets exfiltrated via build logs | tj-actions/changed-files, CVE-2025-30066 (Mar 2025) | Module 2 — SHA-pinning + Build Track L3 secret isolation |
| Distribution/consumption | `curl \| bash` style trust-on-every-run script from vendor infra | Codecov Bash Uploader compromise (2021) | Module 3 |
| Distribution/consumption | Malicious maintainer handoff, backdoor in transitive dependency | event-stream (2018) | Module 3 (context; full fix is the still-Draft Dependencies track) |
| Distribution/consumption | Domain/org takeover of a CDN-hosted include | polyfill.io (2024) | Module 3 (context) |

## Why "stage" != "SLSA track"

SLSA (current spec: v1.2) has two **Approved** (stable) tracks: **Build** and
**Source**. There is no "Distribution track" — the Dependencies and Build
Environment tracks that would eventually cover some of this are still
**Draft**. Distribution/consumption-stage attacks in this table are real and
worth understanding, but they're addressed here as the *verification side*
of Build Track provenance (does anything check it before trusting an
artifact?), not as a leveled SLSA track of their own.
