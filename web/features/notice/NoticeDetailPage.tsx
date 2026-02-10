import { NOTICE_ITEMS } from "./Constants";
import { notFound } from "next/navigation";
import Link from "next/link";
import { ArrowLeft, Calendar } from "lucide-react";

export default function NoticeDetailForm({ id }: { id: string }) {
  const notice = NOTICE_ITEMS.find((item) => item.id === Number(id));

  if (!notice) notFound();

  return (
    <div className="bg-white min-h-screen py-16">
      <div className="container mx-auto px-4 max-w-3xl">
        <Link href="/notice" className="inline-flex items-center text-slate-500 hover:text-blue-700 mb-8 transition">
          <ArrowLeft className="w-4 h-4 mr-1" /> 목록으로 돌아가기
        </Link>

        <article>
          <header className="border-b pb-8 mb-8">
            <span className="text-blue-600 font-bold text-sm uppercase tracking-wider">{notice.category}</span>
            <h1 className="text-3xl md:text-4xl font-bold text-slate-900 mt-2 leading-tight">
              {notice.title}
            </h1>
            <div className="flex items-center gap-2 text-slate-400 mt-4 text-sm">
              <Calendar className="w-4 h-4" /> {notice.date}
            </div>
          </header>

          <div className="text-lg text-slate-700 leading-relaxed whitespace-pre-wrap">
            {notice.content}
          </div>
        </article>
      </div>
    </div>
  );
}