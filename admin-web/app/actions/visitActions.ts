import fs from 'fs/promises';

export async function getVisitorCount() {
  try {
    const VISIT_FILE = process.env.VISIT_FILE;
    const content = await fs.readFile(`${VISIT_FILE}`, 'utf-8');
    const data = JSON.parse(content);

    console.log("Visit Data: ", data);
    
    // 오늘 날짜인지 체크 후 반환
    const today = new Date().toISOString().split('T')[0];
    return data
  } catch (e) {
    return 0;
  }
}