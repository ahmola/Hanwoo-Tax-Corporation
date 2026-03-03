import { NextRequest, NextResponse } from "next/server";

export async function GET(req: NextRequest) {
    const { searchParams} = new URL(req.url);
    const fileUrl = searchParams.get('url');

    console.log("요청 수신 : ", fileUrl);

    if (!fileUrl) return NextResponse.json({error: "Invalid URL"}, {status : 400});

    const title = fileUrl?.split('/').pop() || "download.txt";

    try{
        const res = await fetch(fileUrl!);
        console.log("status: ", res.status);
    
        if (!res.ok) throw new Error(`Fetch failed: ${res.status}`)

        const encodedTitle = encodeURIComponent(title);
        console.log("📂 다운로드 스트리밍 시작: ", encodedTitle);

        return new NextResponse(res.body, {
            headers: {
                'Content-Type': 'application/octet-stream',
                'Content-Disposition': `attachment; filename="${encodedTitle}"; filename*=UTF-8''${encodedTitle}`,
                'Content-Length': res.headers.get('Content-Length') || '',
            },
        })
    }catch(error) {
        console.error("❌ 서버 내부 에러:", error);
        return NextResponse.json({ error: "Internal Server Error" }, { status: 500 });
    }
}