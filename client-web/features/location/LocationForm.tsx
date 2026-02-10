'use client'

import { Map, MapMarker, useKakaoLoader } from "react-kakao-maps-sdk";
import { MapPin, Phone, Clock, Bus, Car } from "lucide-react";

export default function LocationForm() {
  // 1. 카카오맵 로더 사용 (App Router 권장 방식)
  const [loading, error] = useKakaoLoader({
    appkey: process.env.NEXT_PUBLIC_KAKAO_API_KEY as string, // 발급받은 JS 키
    libraries: ["clusterer", "drawing", "services"],
  });

  // 회사 좌표 (한우세무법인 동대문점 실제 좌표로 수정 필요)
  const center = { lat: 37.561845136152115, lng: 127.0143830997867 }; 

  return (
    <section id="location" className="container mx-auto px-4 py-20">
      <div className="flex items-end justify-between mb-8 border-b pb-4 border-slate-200">
        <div>
          <h2 className="head2-with-icon">
            <MapPin className="w-8 h-8 text-blue-700" />
            오시는 길
          </h2>
          <p className="text-slate-500 mt-2">
            한우세무법인 동대문점에 방문해주셔서 감사합니다.
          </p>
        </div>
      </div>

      {/* 2. 레이아웃: 왼쪽(지도) / 오른쪽(주소 정보) */}
      <div className="flex flex-col lg:flex-row gap-8 h-full">
        
        {/* 지도 영역 */}
        <div className="w-full lg:w-2/3 h-[400px] bg-slate-100 rounded-2xl overflow-hidden shadow-md">
          {loading ? (
            // 지도 로딩 중일 때 보여줄 스켈레톤 UI
            <div className="w-full h-full flex items-center justify-center text-slate-400">
              지도를 불러오는 중입니다...
            </div>
          ) : error ? (
            <div className="w-full h-full flex items-center justify-center text-red-500">
              지도를 로드하는데 실패했습니다.
            </div>
          ) : (
            <Map 
              center={center} 
              style={{ width: "100%", height: "100%" }}
              level={3} // 확대 레벨
            >
              <MapMarker position={center}>
                {/* 마커 클릭 시 나타날 정보 (선택사항) */}
                <div className="p-2 text-center text-sm font-bold text-slate-900">
                  한우세무법인 <span className="text-blue-600">동대문점</span>
                </div>
              </MapMarker>
            </Map>
          )}
        </div>

        {/* 주소 및 정보 영역 */}
        <div className="w-full lg:w-1/3 flex flex-col gap-6">
          
          {/* 주소 카드 */}
          <div className="bg-white p-6 rounded-2xl border shadow-sm">
            <h3 className="text-xl font-bold text-slate-900 mb-4 flex items-center gap-2">
              <MapPin className="w-5 h-5 text-blue-600" /> 주소
            </h3>
            <p className="text-slate-600 leading-relaxed">
              서울특별시 중구 다산로 207, 4층<br/>
              (신당동 292-56)
            </p>
          </div>

          {/* 연락처 카드 */}
          <div className="bg-white p-6 rounded-2xl border shadow-sm">
            <h3 className="text-xl font-bold text-slate-900 mb-4 flex items-center gap-2">
              <Phone className="w-5 h-5 text-blue-600" /> 연락처
            </h3>
            <p className="text-slate-600">Tel. 02-2235-4012</p>
            <p className="text-slate-600">Fax. 02-2235-4012</p>
          </div>

          {/* 운영시간 카드 */}
          <div className="bg-white p-6 rounded-2xl border shadow-sm">
            <h3 className="text-xl font-bold text-slate-900 mb-4 flex items-center gap-2">
              <Clock className="w-5 h-5 text-blue-600" /> 운영시간
            </h3>
            <p className="text-slate-600">평일 09:00 - 18:00</p>
            <p className="text-slate-400 text-sm">(점심시간 12:00 - 13:00)</p>
          </div>

        </div>
      </div>
      
      {/* 3. 하단 교통편 안내 (선택사항) */}
      <div className="grid md:grid-cols-2 gap-6 mt-8">
        <div className="flex gap-4 p-5 bg-slate-50 rounded-xl">
          <Bus className="w-10 h-10 text-slate-400 flex-shrink-0" />
          <div>
            <h4 className="font-bold text-slate-800 mb-1">대중교통 이용 시</h4>
            <p className="text-sm text-slate-600">지하철 5, 6호선 청구역 3번 출구 도보 1분</p>
          </div>
        </div>
        <div className="flex gap-4 p-5 bg-slate-50 rounded-xl">
          <Car className="w-10 h-10 text-slate-400 flex-shrink-0" />
          <div>
            <h4 className="font-bold text-slate-800 mb-1">자차 이용 시</h4>
            <p className="text-sm text-slate-600">건물 뒤편 유료 주차장 이용 가능</p>
          </div>
        </div>
      </div>

    </section>
  );
}