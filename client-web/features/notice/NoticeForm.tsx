import Link from "next/link";
import { ChevronRight, Megaphone } from "lucide-react";
import { getNotices } from "./services";

export default async function NoticeForm() {
  const notices = await getNotices();
  const recentNotices = notices.slice(0, 5);

  return (
    <section id="notice" className="container mx-auto px-4">
      {/* 1. 섹션 헤더: 제목과 '더보기' 버튼 */}
      <div className="flex items-end justify-between mb-8 border-b pb-4 border-slate-200">
        <div>
          <h2 className="head2-with-icon">
            <Megaphone className="w-8 h-8 text-blue-700" />
            공지사항
          </h2>
          <p className="text-slate-500 mt-2">
            한우세무법인 동대문점의 새로운 소식과 주요 세무 일정을 알려드립니다.
          </p>
        </div>
        <Link 
          href="/notice" // 공지사항 전체보기 페이지로 연결
          className="hidden md:flex items-center text-sm font-semibold text-slate-500 hover:text-blue-700 transition"
        >
          전체보기 <ChevronRight className="w-4 h-4 ml-1" />
        </Link>
      </div>

      {/* 2. 공지사항 리스트 */}
      <div className="flex flex-col gap-4">
        {recentNotices.length > 0 ? (
          recentNotices.map((item) => (
            <Link
              key={item.id}
              href={`/notice/${item.id}`} // 개별 상세 페이지 링크 (추후 구현)
              className="group flex flex-col md:flex-row md:items-center justify-between p-5 bg-white border rounded-xl hover:border-blue-300 hover:shadow-md transition duration-200"
            >
              <div className="flex flex-col md:flex-row md:items-center gap-2 md:gap-6 flex-1">
                {/* 날짜 & 카테고리 뱃지 */}
                <div className="flex items-center gap-3">
                  <span className={`px-2 py-1 text-xs font-bold rounded ${
                    item.isImportant 
                      ? "bg-red-50 text-red-600" 
                      : "bg-slate-100 text-slate-600"
                  }`}>
                    {item.category}
                  </span>
                  <span className="text-sm text-slate-400 font-medium">
                    {item.date}
                  </span>
                </div>
                
                {/* 제목 */}
                <h3 className="text-lg font-medium text-slate-800 group-hover:text-blue-700 transition line-clamp-1">
                  {item.title}
                </h3>
              </div>

              {/* 화살표 아이콘 (호버시 이동 효과) */}
              <div className="hidden md:block text-slate-300 group-hover:text-blue-500 group-hover:translate-x-1 transition flex-shrink-0 ml-4">
                <ChevronRight />
              </div>
            </Link>
          ))
        ) : (
            <div className="text-center py-10 text-slate-400 bg-slate-50 rounded-xl">
                등록된 공지사항이 없습니다.
            </div>
          )
        }
      </div>

      {/* 모바일용 더보기 버튼 (모바일에서만 하단에 표시) */}
      <div className="mt-6 md:hidden text-center">
        <Link 
          href="/notice" 
          className="inline-flex items-center text-sm font-semibold text-slate-500 hover:text-blue-700"
        >
          전체 소식 보기 <ChevronRight className="w-4 h-4 ml-1" />
        </Link>
      </div>
    </section>
  );
}