'use server'
import fs from 'fs/promises';
import path from 'path';

const VISIT_FILE = process.env.VISIT_FILE || '/usr/share/next-resources/visitor_count.json'

export async function countVisitor() {
    try{
        // 초기 데이터 설정
        const today = new Date().toISOString().split('T')[0];
        let data = {date: today, count: 0};

        // 디렉터리가 없다면 만듦
        const dir = path.dirname(`${VISIT_FILE}`);
        await fs.mkdir(dir, {recursive:true});

        // 파일 존재 여부 확인, 없다면 생성
        try {
            await fs.access(`${VISIT_FILE}`);
        }catch {
            const initialData = JSON.stringify(data);
            console.log("방문자 카운트 파일 생성: ", initialData)
            // 파일 생성
            await fs.writeFile(`${VISIT_FILE}`, initialData);

        }

        // json 파일 읽음
        const content = await fs.readFile(`${VISIT_FILE}`, 'utf-8');
        data = JSON.parse(content);
        console.log("방문자 카운트 파일 읽음: ", data)

        // 매일 초기화
        if (data.date !== today) {
            data = {date: today, count: 1};
        } else{
            data.count += 1;
        }

        console.log("방문자 카운트: ", data);

        // json 파일 업데이트
        try {
            await fs.writeFile(`${VISIT_FILE}`, JSON.stringify(data));
        } catch(error){
            console.log("방문자 카운트 업데이트 에러: ", error)
        }
    }catch(error) {
        console.log("방문자 카운트 에러: ", error)
    }
}