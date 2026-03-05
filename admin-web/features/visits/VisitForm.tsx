export const dynamic = 'force-dynamic';

import { getVisitorCount } from "@/app/actions/visitActions";

export default async function VisitForm() {
  const count = await getVisitorCount();
  console.log("Count: ",count);
  return (
    <div className="bg-white p-6 rounded-xl shadow-sm border border-slate-200">
        <h3 className="text-slate-500 text-sm font-semibold">오늘 방문자수</h3>
        <p className="text-3xl font-bold text-slate-800 mt-2">{count}명</p>
    </div>
  );
}