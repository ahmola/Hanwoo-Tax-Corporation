'use client' // 검색 기능 등 인터랙션을 위해 클라이언트 컴포넌트로 지정

import Link from "next/link";
import { Search, ChevronRight, Megaphone, Calendar, Tag } from "lucide-react";
import { getNotices } from "@/app/actions/noticeActions";
import { useEffect, useState } from "react";
import { NoticeItem } from "./Constants";

export default function NoticePage() {
  const [searchTerm, setSearchTerm] = useState("");
  const [notices, setNotices] = useState<NoticeItem[]>([]);

  useEffect(() => { 
    getNotices().then(res => setNotices(res));
  }, []);

  // 검색어 필터링 로직 (제목으로 검색)
  const filteredItems = notices.filter((item) =>
    item.title.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <section className="container mx-auto px-4 py-20 min-h-screen">
      
      {/* 1. 헤더 영역 */}
      <div className="text-center mb-12">
        <h2 className="text-3xl font-bold text-white-900 flex justify-center items-center gap-2">
          <Megaphone className="w-8 h-8 text-blue-700" />
          공지사항 전체보기
        </h2>
        <p className="text-slate-500 mt-2">
          세무법인 넥스트의 모든 소식을 확인하세요.
        </p>
      </div>

      {/* 2. 검색창 및 상단 툴바 */}
      <div className="max-w-4xl mx-auto mb-8 flex justify-between items-center bg-white p-4 rounded-xl shadow-sm border">
        <div className="relative w-full md:w-96">
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400 w-5 h-5" />
          <input 
            type="text"
            placeholder="검색어를 입력하세요..." 
            className="input-primary"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>
        <div className="hidden md:block text-slate-500 text-sm">
          총 <span className="font-bold text-blue-700">{filteredItems.length}</span>건
        </div>
      </div>

      {/* 3. 리스트 영역 */}
      <div className="max-w-4xl mx-auto flex flex-col gap-4">
        {filteredItems.length > 0 ? (
          filteredItems.map((item) => (
            <Link
              key={item.id}
              href={`/notice/${item.id}`} // 상세 페이지로 이동
              className="group bg-white border rounded-xl p-6 hover:shadow-lg hover:border-blue-400 transition duration-200 relative overflow-hidden"
            >
              {/* 왼쪽 중요 표시 줄 */}
              {item.isImportant && (
                <div className="absolute left-0 top-0 bottom-0 w-1.5 bg-red-500" />
              )}

              <div className="flex flex-col md:flex-row md:items-center justify-between gap-4">
                <div className="space-y-2">
                  <div className="flex items-center gap-2">
                    {/* 카테고리 뱃지 */}
                    <span className={`px-2 py-1 text-xs font-bold rounded flex items-center gap-1 ${
                      item.isImportant ? "bg-red-50 text-red-600" : "bg-slate-100 text-slate-600"
                    }`}>
                      <Tag className="w-3 h-3" /> {item.category}
                    </span>
                    {/* 날짜 */}
                    <span className="text-sm text-slate-400 flex items-center gap-1">
                      <Calendar className="w-3 h-3" /> {item.date}
                    </span>
                  </div>
                  
                  {/* 제목 */}
                  <h3 className="text-lg md:text-xl font-bold text-slate-800 group-hover:text-blue-700 transition">
                    {item.title}
                  </h3>
                </div>

                {/* 화살표 아이콘 */}
                <div className="hidden md:flex items-center text-slate-300 group-hover:text-blue-600 group-hover:translate-x-1 transition">
                  <ChevronRight />
                </div>
              </div>
            </Link>
          ))
        ) : (
          /* 검색 결과 없음 */
          <div className="text-center py-20 text-slate-500">
            검색 결과가 없습니다.
          </div>
        )}
      </div>

    </section>
  );
}