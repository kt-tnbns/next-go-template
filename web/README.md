# Next.js Frontend

A modern Next.js frontend application with TypeScript, Tailwind CSS, and seamless integration with the NestJS backend.

## 🚀 Getting Started

### Prerequisites

- Node.js >= 20.x
- pnpm (recommended) or npm/yarn

### Installation

```bash
# Install dependencies
pnpm install

# Set up environment variables
cp .env.template .env.local

# Start development server
pnpm dev
```

Open [http://localhost:3001](http://localhost:3001) with your browser to see the result.

## 📁 Project Structure

```
src/
├── app/               # Next.js App Router
│   ├── services/      # API service modules
│   ├── types/         # TypeScript type definitions
│   ├── page.tsx       # Home page
│   └── layout.tsx     # Root layout
└── lib/               # Shared libraries
    ├── api.ts         # Axios client configuration
    └── services.ts    # API service functions
```

## 🔌 API Integration

### Configuration

The API client is configured in `src/lib/api.ts`:

```typescript
import axios from 'axios';

export const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});
```

### Environment Variables

Configure in `.env.local`:

```env
# Backend API URL
NEXT_PUBLIC_API_URL=http://localhost:3000

# Frontend Port
PORT=3001
```

### Creating API Services

Create typed service functions in `src/lib/services.ts`:

```typescript
import { api } from '@/lib/api';

export interface User {
  id: number;
  email: string;
  name: string;
}

export const userService = {
  getAll: () => api.get<{ success: boolean; data: User[] }>('/api/users'),
  getById: (id: number) => api.get<{ success: boolean; data: User }>(`/api/users/${id}`),
  create: (data: Omit<User, 'id'>) =>
    api.post<{ success: boolean; data: User }>('/api/users', data),
  update: (id: number, data: Partial<User>) =>
    api.patch<{ success: boolean; data: User }>(`/api/users/${id}`, data),
  delete: (id: number) => api.delete<{ success: boolean }>(`/api/users/${id}`),
};
```

### Using Services in Components

```typescript
'use client';

import { useEffect, useState } from 'react';
import { userService, type User } from '@/lib/services';

export default function UsersPage() {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const { data } = await userService.getAll();
        if (data.success) {
          setUsers(data.data);
        }
      } catch (error) {
        console.error('Failed to fetch users:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchUsers();
  }, []);

  if (loading) return <div>Loading...</div>;

  return (
    <div>
      <h1>Users</h1>
      <ul>
        {users.map((user) => (
          <li key={user.id}>{user.name}</li>
        ))}
      </ul>
    </div>
  );
}
```

## 🎨 Styling

This project uses **Tailwind CSS 4** for styling.

### Utility Classes

```tsx
<div className="flex items-center justify-center p-4 bg-blue-500 text-white rounded-lg">
  Hello, World!
</div>
```

### Custom Styles

For component-specific styles, use the `style` attribute or CSS modules:

```tsx
// Inline styles
<div style={{ fontSize: '16px', color: '#333' }}>Custom style</div>

// Or use Tailwind utilities
<div className="text-base text-gray-800">Same effect</div>
```

## 🔧 Available Scripts

```bash
# Development
pnpm dev              # Start development server
pnpm build            # Build for production
pnpm start            # Start production server

# Code Quality
pnpm lint             # Run ESLint
pnpm format           # Format code with Prettier
```

## 📖 Conventions

### File Organization

- **Pages**: `src/app/page.tsx`, `src/app/about/page.tsx`
- **Layouts**: `src/app/layout.tsx`, `src/app/about/layout.tsx`
- **Components**: Co-located with pages or in `src/components/`
- **Services**: `src/lib/services.ts` or `src/app/services/`
- **Types**: `src/lib/types.ts` or `src/app/types/`

### Naming Conventions

- **Components**: PascalCase (`UserProfile.tsx`)
- **Utilities**: camelCase (`formatDate.ts`)
- **Types**: PascalCase (`interface User {}`)
- **Constants**: UPPER_SNAKE_CASE (`const API_URL`)

### Code Style

- Use TypeScript for type safety
- Use functional components with hooks
- Keep components small and focused
- Use Tailwind CSS for styling (avoid inline styles when possible)
- Define types for API responses
- Handle errors gracefully
- Use `use client` directive for client components (interactivity)

## 🌐 Routing

This project uses the **Next.js App Router** (App Directory).

### File-Based Routing

```
app/
├── page.tsx              # → /
├── about/
│   └── page.tsx          # → /about
├── users/
│   ├── page.tsx          # → /users
│   └── [id]/
│       └── page.tsx      # → /users/123
└── layout.tsx            # Root layout (applies to all pages)
```

### Dynamic Routes

```tsx
// app/users/[id]/page.tsx
export default function UserPage({ params }: { params: { id: string } }) {
  return <div>User ID: {params.id}</div>;
}
```

## 📦 Adding a New Page

1. Create a new directory in `app/`:

```bash
mkdir -p src/app/about
```

2. Create `page.tsx`:

```tsx
export default function AboutPage() {
  return (
    <div>
      <h1>About</h1>
      <p>This is the about page.</p>
    </div>
  );
}
```

3. Optionally create a `layout.tsx` for page-specific layout:

```tsx
export default function AboutLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="about-layout">
      <nav>About Navigation</nav>
      {children}
    </div>
  );
}
```

## 🔐 Authentication (Future)

To enable authentication:

1. Set `withCredentials: true` in `src/lib/api.ts`:

```typescript
export const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  withCredentials: true, // Enable cookies
});
```

2. Update backend CORS settings:

```env
# core/.env
CORS_ALLOWED_ORIGINS=http://localhost:3001
CORS_CREDENTIALS=true
```

3. Implement auth pages in `app/auth/` directory.

## 🧪 Testing (Future)

Testing setup can be added with:

```bash
pnpm add -D @testing-library/react @testing-library/jest-dom jest-environment-jsdom
```

## 📚 Learn More

- [Next.js Documentation](https://nextjs.org/docs) - Learn about Next.js features
- [Tailwind CSS Documentation](https://tailwindcss.com/docs) - Learn about utility classes
- [TypeScript Documentation](https://www.typescriptlang.org/docs/) - Learn about TypeScript

## 🚀 Deployment

### Vercel (Recommended)

```bash
# Install Vercel CLI
pnpm add -g vercel

# Deploy
vercel
```

### Docker

```bash
# Build image
docker build -t nextjs-frontend .

# Run container
docker run -p 3001:3001 nextjs-frontend
```

### Environment Variables

Ensure `NEXT_PUBLIC_API_URL` is set in your deployment environment.

## 🤝 Contributing

1. Follow the code style conventions
2. Use TypeScript for type safety
3. Test your changes before committing
4. Run linter and formatter:
   ```bash
   pnpm lint
   pnpm format
   ```

## 📝 License

MIT
