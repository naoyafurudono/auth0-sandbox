"use client";

import { useUser } from "@auth0/nextjs-auth0/client";

export default function Home() {
  const { user, error, isLoading } = useUser();

  if (isLoading) return <div className="p-8">読み込み中...</div>;
  if (error) return <div className="p-8">エラー: {error.message}</div>;

  return (
    <main className="min-h-screen p-8">
      <div className="max-w-4xl mx-auto">
        <h1 className="text-4xl font-bold mb-8">Auth0 Sandbox</h1>

        {user ? (
          <div className="space-y-4">
            <div className="p-4 bg-green-50 border border-green-200 rounded">
              <p className="text-green-800">ログイン中: {user.name}</p>
              <p className="text-sm text-green-600">{user.email}</p>
            </div>
            <div className="space-x-4">
              <a
                href="/profile"
                className="inline-block px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
              >
                プロフィール
              </a>
              <a
                href="/api/auth/logout"
                className="inline-block px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600"
              >
                ログアウト
              </a>
            </div>
          </div>
        ) : (
          <div>
            <p className="mb-4">Auth0を使用したOIDC認証のデモです。</p>
            <a
              href="/api/auth/login"
              className="inline-block px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
            >
              ログイン
            </a>
          </div>
        )}
      </div>
    </main>
  );
}
