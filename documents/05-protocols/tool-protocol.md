# Tool Protocol

Version: 1.0

---

## Philosophy

The AI never accesses business systems directly.

All business capabilities are exposed as tools.

---

## Tool Lifecycle

Plan

↓

Invoke Tool

↓

Observe Result

↓

Reason

↓

Generate Response

---

## Tool Categories

Discovery

Recommendation

Comparison

Knowledge

Commerce

Memory

Analytics

---

## Tool Contract

Every tool exposes

Name

Description

Input Schema

Output Schema

Error Schema

---

## Tool Response

Success

Data

Metadata

Observations

---

## Error Types

Validation

Permission

Provider Failure

Timeout

Not Found

Business Error

---

## Rules

Tools are deterministic.

Tools never return natural language.

Tools return structured JSON only.

The LLM performs reasoning.

Business logic belongs to backend.
