import Link from "next/link";
import { FileText, Download, ChevronRight, FolderOpen } from "lucide-react";
import { getDocuments } from "@/app/actions/documentActions";

export default async function DocumentsForm() {
  const documents = await getDocuments();
  const recentDocuments = documents.slice(0, 5);

  return (
    <section id="documents" className="container mx-auto px-4">
      {/* 헤더 */}
      <div className="flex items-end justify-between mb-8 border-b pb-4 border-slate-200">
        <div>
          <h2 className="head2-with-icon">
            <FolderOpen className="w-8 h-8 text-blue-700" />
            문서
          </h2>
          <p className="text-slate-500 mt-2">
            세무 업무에 필요한 필수 서식과 안내문을 다운로드하세요.
          </p>
        </div>
        <Link 
          href="/documents" 
          className="hidden md:flex items-center text-sm font-semibold text-slate-500 hover:text-blue-700 transition"
        >
          문서 더보기 <ChevronRight className="w-4 h-4 ml-1" />
        </Link>
      </div>

      {/* 리스트 (최근 4개) */}
      <div className="grid md:grid-cols-2 gap-4">
        {recentDocuments.length > 0 ? (
          documents.map((item) => (
            <div
              key={item.id}
              className="group flex items-center justify-between p-5 bg-white border rounded-xl hover:border-blue-300 hover:shadow-md transition duration-200"
            >
              <div className="flex items-center gap-4 overflow-hidden">
                <div className="p-3 bg-slate-50 rounded-lg group-hover:bg-blue-50 transition">
                  <FileText className="w-6 h-6 text-slate-400 group-hover:text-blue-600" />
                </div>
                <div className="min-w-0">
                  <h3 className="font-bold text-slate-800 truncate group-hover:text-blue-700 transition">
                    {item.title}
                  </h3>
                  <p className="text-sm text-slate-400 mt-1">
                    {item.category} · {item.date}
                  </p>
                </div>
              </div>

              {/* 다운로드 버튼 (실제 다운로드는 a 태그 사용) */}
              <a 
                href={item.fileUrl} 
                download 
                className="p-2 text-slate-400 hover:text-blue-700 hover:bg-slate-100 rounded-full transition ml-2 flex-shrink-0"
                title="다운로드"
              >
                <Download className="w-5 h-5" />
              </a>
            </div>
          ))
          ) : (
            <div className="text-center py-10 text-slate-400 bg-slate-50 rounded-xl">
              등록된 공지사항이 없습니다.
            </div>
          )
        }
      </div>
      
      {/* 모바일 더보기 버튼 */}
      <div className="mt-6 md:hidden text-center">
        <Link 
          href="/documents" 
          className="inline-flex items-center text-sm font-semibold text-slate-500 hover:text-blue-700"
        >
          자료실 전체보기 <ChevronRight className="w-4 h-4 ml-1" />
        </Link>
      </div>
    </section>
  );
}