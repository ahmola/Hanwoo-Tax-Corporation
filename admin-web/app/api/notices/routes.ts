import { NextResponse } from "next/server";

// GET: 목록 가져오기 테스트용
export async function GET() {
  return NextResponse.json([
    { id: 1, title: "가짜 공지사항입니다 1", category: "공지", date: "2026-02-11" },
    { id: 2, title: "가짜 공지사항입니다 2", category: "세무일정", date: "2026-02-12" },
  ]);
}

// POST: 등록 테스트용
export async function POST(request: Request) {
  const data = await request.json();
  console.log("받은 데이터:", data);
  return NextResponse.json({ success: true, message: "임시 저장 완료" });
}