// 드롭다운 메뉴
'use client'

import Link from "next/link";
import { NAV_ITEMS } from "@/config/navigation";

interface MenuProps{
    isOpen: boolean;
    onClose: () => void;
}

export default function DropDownMenu({isOpen, onClose}: MenuProps) {
    if (!isOpen) return null;

    return (
        <div className="absolute top-16 left-0 w-full bg-white border-b shadow-xl animate-in slide-in-from-top-2 z-40">
            <div className="container mx-auto flex flex-col p-4 gap-2">
                {NAV_ITEMS.map((item) => (
                    <Link
                        key={item.href}
                        href={item.href}
                        className="p-3 text-slate-700 hover:bg-slate-50 hover:text-blue-700 rounded-lg transition font-medium"
                        onClick={onClose}
                    >
                        {item.label}
                    </Link>
                ))}
            </div>
        </div>
  );
}