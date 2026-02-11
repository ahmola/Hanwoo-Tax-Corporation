import { NextResponse } from "next/server";

export async function GET() {
  return NextResponse.json([
    { id: 1, name: "테스터", phone: "010-1234-5678", message: "상담 신청합니다.", createdAt: "2026-02-11" },
  ]);
}