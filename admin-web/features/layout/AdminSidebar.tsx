// src/features/layout/AdminSidebar.tsx
import Link from "next/link";
import { FileText, Users, Megaphone, FolderOpen } from "lucide-react";

export default function AdminSidebar() {
  return (
    <aside className="w-64 bg-slate-900 text-white min-h-screen p-6 fixed left-0 top-0">
      <Link href="/">
        <h1 className="text-2xl font-bold mb-10">관리자 페이지</h1>
      </Link>
      <nav className="flex flex-col gap-2">
        <Link href="/notices/create" className="menu-text">
          <Megaphone size={20} /> 공지사항 등록
        </Link>
        <Link href="/documents/upload" className="menu-text">
          <FolderOpen size={20} /> 자료실 업로드
        </Link>
        <Link href="/contacts" className="menu-text">
          <Users size={20} /> 상담 신청 내역
        </Link>
      </nav>
    </aside>
  );
}