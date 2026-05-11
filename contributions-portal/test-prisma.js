require('dotenv').config({ path: '.env.local' });
const { PrismaClient } = require('@prisma/client');
const prisma = new PrismaClient();
prisma.submission.findMany().then(console.log).catch(console.error).finally(() => prisma.$disconnect());
