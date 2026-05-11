const { PrismaClient } = require('@prisma/client');
const prisma = new PrismaClient();

async function main() {
  console.log('--- Subs ---');
  console.log(await prisma.$queryRaw`SELECT id, contributor_id, analysis_result_id FROM html_css_submissions`);
  console.log('--- Results ---');
  console.log(await prisma.$queryRaw`SELECT id, contributor_id, status FROM analysis_results`);
}

main()
  .catch(console.error)
  .finally(() => prisma.$disconnect());
