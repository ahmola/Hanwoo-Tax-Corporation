// 레이아웃 헤더
'use client';

import Link from "next/link";
import { ShieldCheck, Menu, X } from "lucide-react";
import { useState } from "react";
import DropDownMenu from "@/components/layout/DropDownMenu";

export default function Header() {
    const [isMenuOpen, setIsMenuOpen] = useState(false);
    
    return (
    <header className="fixed top-0 w-full bg-white/80 backdrop-blur-md border-b z-50">
        <div className="container mx-auto px-4 h-16 flex items-center justify-between">
            {/* 로고 */}
            <Link href="/" className="flex items-center gap-2 font-bold text-xl text-slate-900">
                <ShieldCheck className="text-blue-700" />
                <span>한우세무법인 동대문점</span>
            </Link>

            {/* 우측 버튼 */}
            <div className="flex items-center gap-3">
                {/* 상담 신청 버튼 (항상 보임) */}
                <Link 
                    href="#contact" 
                    className="px-4 py-2 bg-blue-700 text-white text-sm font-bold rounded-md hover:bg-blue-800 transition"
                >
                    상담 신청
                </Link>

                {/* 햄버거 메뉴 버튼 (항상 보임) */}
                <button 
                    className="p-2 text-slate-600 hover:text-slate-900 transition hover:bg-slate-100 rounded-md"
                    onClick={() => setIsMenuOpen(!isMenuOpen)}
                    aria-label={isMenuOpen ? "메뉴 열기" : "메뉴 열기"}
                >
                    {isMenuOpen ? <X size={24} /> : <Menu size={24} />} 
                </button>
            </div>
        </div>

        {/* 3. 드롭다운 메뉴 (isMenuOpen이 true일 때만 열림) */}
        <DropDownMenu isOpen={isMenuOpen} onClose={() => setIsMenuOpen(false)} />
    </header>
  );
}