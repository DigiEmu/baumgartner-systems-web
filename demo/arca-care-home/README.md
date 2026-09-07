# Arca Care-Home Demo

This package is a fictional dataset for demonstrating Arca to a care-home administrator or manager. It contains no real resident information.

## Objective
Show that an AI-assisted administrative workflow can be accompanied by a verifiable evidence trail.

## Files
- `resident-request.txt`
- `policy.md`
- `staff-note.txt`
- `expected-output.md`

## Demo flow

### 1. Import
Import the three source documents and show that Arca records the evidence basis and source identity.

### 2. Build
Create the administrative result. The expected content is in `expected-output.md`.

### 3. Verify
Verify the recorded case. Optionally change one sentence in `staff-note.txt` and verify again to demonstrate that changed evidence is detectable.

### 4. Reconstruct
Return to the original evidence set and reconstruct the case.

## 3–5 minute sales narrative
1. Here is a small administrative request.
2. These are the exact documents available to the workflow.
3. Arca records their identity and provenance.
4. We produce the resulting administrative record.
5. We verify that the evidence has not silently changed.
6. Later we can reconstruct what happened from the recorded evidence.

## Scope
This demo shows evidence handling, provenance, verification and reconstruction concepts. It does not validate medical decisions, legal correctness, policy authority or semantic truth beyond the supplied demonstration evidence.
