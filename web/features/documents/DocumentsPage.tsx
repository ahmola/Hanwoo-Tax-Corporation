'use client'

import { useState } from "react";
import Link from "next/link";
import { DOCUMENT_ITEMS } from "@/features/documents/Constants";
import { Search, FolderOpen, ArrowLeft, Download, FileText } from "lucide-react";

export default function DocumentsPage() {
  const [searchTerm, setSearchTerm] = useState("");

  // 검색 필터링
  const filteredItems = DOCUMENT_ITEMS.filter((item) =>
    item.title.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div className="bg-slate-50 min-h-screen py-20">
      <div className="container mx-auto px-4 max-w-4xl">
        
        {/* 뒤로가기 */}
        <Link href="/" className="inline-flex items-center text-slate-500 hover:text-blue-700 mb-8 font-medium">
          <ArrowLeft className="w-4 h-4 mr-2" /> 메인으로 돌아가기
        </Link>

        {/* 헤더 & 검색창 */}
        <div className="bg-white p-8 rounded-2xl border shadow-sm mb-8">
          <div className="text-center mb-8">
            <h1 className="text-3xl font-bold text-slate-900 flex justify-center items-center gap-2">
              <FolderOpen className="w-8 h-8 text-blue-700" />
              자료실
            </h1>
            <p className="text-slate-500 mt-2">필요한 서식을 검색해서 다운로드하세요.</p>
          </div>

          <div className="relative max-w-lg mx-auto">
            <Search className="absolute left-4 top-1/2 -translate-y-1/2 text-slate-400 w-5 h-5" />
            <input 
              type="text"
              placeholder="자료 제목을 검색하세요..." 
              className="input-primary"
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
            />
          </div>
        </div>

        {/* 리스트 */}
        <div className="flex flex-col gap-3">
          {filteredItems.length > 0 ? (
            filteredItems.map((item) => (
              <div
                key={item.id}
                className="group flex items-center justify-between p-5 bg-white border rounded-xl hover:border-blue-400 hover:shadow-md transition duration-200"
              >
                <div className="flex items-center gap-4 overflow-hidden">
                  <div className="w-12 h-12 flex items-center justify-center bg-blue-50 text-blue-600 rounded-lg flex-shrink-0">
                    <FileText className="w-6 h-6" />
                  </div>
                  <div className="min-w-0">
                    <div className="flex items-center gap-2 mb-1">
                      <span className="text-xs font-bold px-2 py-0.5 bg-slate-100 text-slate-600 rounded">
                        {item.category}
                      </span>
                      <span className="text-xs text-slate-400">{item.size}</span>
                    </div>
                    <h3 className="text-lg font-bold text-slate-800 truncate group-hover:text-blue-700 transition">
                      {item.title}
                    </h3>
                  </div>
                </div>

                {/* 다운로드 버튼 */}
                <a 
                  href={item.fileUrl}
                  download
                  className="flex items-center gap-2 px-4 py-2 bg-slate-100 text-slate-600 rounded-lg hover:bg-blue-600 hover:text-white transition font-semibold text-sm flex-shrink-0 ml-4"
                >
                  <Download className="w-4 h-4" />
                  <span className="hidden md:inline">다운로드</span>
                </a>
              </div>
            ))
          ) : (
            <div className="text-center py-20 text-slate-500">
              검색 결과가 없습니다.
            </div>
          )}
        </div>

      </div>
    </div>
  );
}