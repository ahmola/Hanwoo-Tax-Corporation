import type { Metadata } from "next";
import "./globals.css";
import AdminSidebar from "@/features/layout/AdminSidebar";

export const metadata: Metadata = {
  title: "관리자 페이지",
  description: "한우세무법인 동대문점 관리자 페이지",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ko">
      <body className="flex min-h-screen bg-slate-100 text-slate-900">
        {/* 1. 사이드바 (왼쪽 고정) */}
        <AdminSidebar />
        
        {/* 2. 메인 콘텐츠 영역 (오른쪽 나머지) */}
        <main className="flex-1 ml-64 p-8 overflow-y-auto">
          {children}
        </main>
      </body>
    </html>
  );
}