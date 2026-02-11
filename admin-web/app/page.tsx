export default function AdminDashboard() {
  return (
    <div>
      <h1 className="text-3xl font-bold mb-8">대시보드</h1>
      
      {/* 요약 카드 그리드 */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
        <div className="bg-white p-6 rounded-xl shadow-sm border border-slate-200">
          <h3 className="text-slate-500 text-sm font-semibold">신규 상담 신청</h3>
          <p className="text-3xl font-bold text-blue-600 mt-2">12건</p>
        </div>
        <div className="bg-white p-6 rounded-xl shadow-sm border border-slate-200">
          <h3 className="text-slate-500 text-sm font-semibold">오늘의 방문자</h3>
          <p className="text-3xl font-bold text-slate-800 mt-2">1,240명</p>
        </div>
        <div className="bg-white p-6 rounded-xl shadow-sm border border-slate-200">
          <h3 className="text-slate-500 text-sm font-semibold">서버 상태</h3>
          <p className="text-3xl font-bold text-green-600 mt-2">정상</p>
        </div>
      </div>

      <div className="bg-white p-8 rounded-xl shadow-sm border border-slate-200 text-center py-20 text-slate-400">
        최근 활동 내역이 여기에 표시됩니다.
      </div>
    </div>
  );
}