// 공통 페이지 툴, 기본 레이아웃
import type { Metadata } from "next";
import {Inter} from "next/font/google";
import "./globals.css";
import Header from "@/components/layout/Header";

const inter = Inter({ subsets: ["latin"] })

export const metadata: Metadata = {
  title: "한우세무법인 동대문점",
  description: "신뢰할 수 있는 세무 파트너, 한우세무법인",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="ko">
      <body className={inter.className}>
        <Header />
        <main className="min-h-screen pt-16">{children}</main>
        <footer className="bg-slate-900 text-white py-8 text-center text-sm">
          © 2026 한우세무법인 동대문점. All rights reserved.
        </footer>
      </body>
    </html>
  );
}
