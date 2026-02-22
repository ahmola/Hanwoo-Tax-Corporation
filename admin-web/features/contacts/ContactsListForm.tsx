'use client'

import { useEffect, useState } from "react";
import { getContacts } from "@/app/actions/contactActions";

interface Contact {
  id: number;
  name: string;
  phone: string;
  message: string;
  created_at: string;
}

export default function ContactsList() {
  const [list, setList] = useState<Contact[]>([]);

  // 페이지 로드 시 API 서버에서 목록 가져오기
  useEffect(() => {
    const fetchList = async () => {
      const data = await getContacts();

      setList(data);
    };

    fetchList().then(data => {
      console.log(data);
    });
  }, []);

  return (
    <div className="bg-white p-6 rounded-xl border">
      <h2 className="table-head-name">상담 신청 내역</h2>
      <table className="w-full text-left border-collapse">
        <thead>
          <tr className="border-b bg-slate-50 text-slate-500">
            <th className="p-3">날짜</th>
            <th className="p-3">이름</th>
            <th className="p-3">연락처</th>
            <th className="p-3">내용</th>
          </tr>
        </thead>
        <tbody>
          {list.map((item) => (
            <tr key={item.id} className="border-b hover:bg-slate-50">
              <td className="date-primary">{item.created_at}</td>
              <td className="table-element">{item.name}</td>
              <td className="table-element">{item.phone}</td>
              <td className="table-text">{item.message}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}