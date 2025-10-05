"use client";

import { useUser } from "@auth0/nextjs-auth0/client";
import { zodResolver } from "@hookform/resolvers/zod";
import { useId } from "react";
import { useForm } from "react-hook-form";
import useSWR from "swr";
import { z } from "zod";
import { fetcher } from "@/lib/fetcher";

interface Profile {
  displayName?: string;
  bio?: string;
}

interface UserData {
  id: string;
  content: string;
  createdAt: string;
}

const profileSchema = z.object({
  displayName: z.string(),
  bio: z.string(),
});

const dataSchema = z.object({
  content: z.string().min(1, "コンテンツを入力してください"),
});

type ProfileForm = z.infer<typeof profileSchema>;
type DataForm = z.infer<typeof dataSchema>;

export default function ProfilePage() {
  const { user, isLoading } = useUser();

  const displayNameId = useId();
  const bioId = useId();
  const newContentId = useId();

  const { data: profile, mutate: mutateProfile } = useSWR<Profile>(
    user ? "/api/backend/api/v1/users/me/profile" : null,
    fetcher,
  );

  const { data: userData = [], mutate: mutateUserData } = useSWR<UserData[]>(
    user ? "/api/backend/api/v1/users/me/data" : null,
    fetcher,
  );

  const { register: registerProfile, handleSubmit: handleSubmitProfile } =
    useForm<ProfileForm>({
      resolver: zodResolver(profileSchema),
      values: {
        displayName: profile?.displayName ?? "",
        bio: profile?.bio ?? "",
      },
    });

  const {
    register: registerData,
    handleSubmit: handleSubmitData,
    reset: resetData,
    formState: { errors },
  } = useForm<DataForm>({
    resolver: zodResolver(dataSchema),
    defaultValues: {
      content: "",
    },
  });

  const updateProfile = async (data: ProfileForm) => {
    try {
      const response = await fetch("/api/backend/api/v1/users/me/profile", {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
      });
      if (response.ok) {
        alert("プロフィール更新成功！");
        mutateProfile();
      }
    } catch (error) {
      console.error("Failed to update profile:", error);
      alert("プロフィール更新失敗");
    }
  };

  const createData = async (data: DataForm) => {
    try {
      const response = await fetch("/api/backend/api/v1/users/me/data", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
      });
      if (response.ok) {
        resetData();
        mutateUserData();
      }
    } catch (error) {
      console.error("Failed to create data:", error);
    }
  };

  if (isLoading) return <div className="p-8">読み込み中...</div>;
  if (!user) return <div className="p-8">ログインが必要です</div>;

  return (
    <main className="min-h-screen p-8">
      <div className="max-w-4xl mx-auto">
        <h1 className="text-4xl font-bold mb-8">プロフィール</h1>

        <div className="mb-8 p-4 bg-gray-50 border rounded">
          <h2 className="text-xl font-bold mb-2">Auth0ユーザー情報</h2>
          <p>名前: {user.name}</p>
          <p>メール: {user.email}</p>
        </div>

        <div className="mb-8 p-4 bg-white border rounded">
          <h2 className="text-xl font-bold mb-4">プロフィール編集</h2>
          <form
            onSubmit={handleSubmitProfile(updateProfile)}
            className="space-y-4"
          >
            <div>
              <label htmlFor={displayNameId} className="block mb-1">
                表示名
              </label>
              <input
                id={displayNameId}
                type="text"
                {...registerProfile("displayName")}
                className="w-full border rounded px-3 py-2"
              />
            </div>
            <div>
              <label htmlFor={bioId} className="block mb-1">
                自己紹介
              </label>
              <textarea
                id={bioId}
                {...registerProfile("bio")}
                className="w-full border rounded px-3 py-2"
                rows={4}
              />
            </div>
            <button
              type="submit"
              className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
            >
              更新
            </button>
          </form>
        </div>

        <div className="mb-8 p-4 bg-white border rounded">
          <h2 className="text-xl font-bold mb-4">データ作成</h2>
          <form onSubmit={handleSubmitData(createData)} className="space-y-4">
            <div>
              <label htmlFor={newContentId} className="block mb-1">
                コンテンツ
              </label>
              <input
                id={newContentId}
                type="text"
                {...registerData("content")}
                className="w-full border rounded px-3 py-2"
              />
              {errors.content && (
                <p className="text-red-500 text-sm mt-1">
                  {errors.content.message}
                </p>
              )}
            </div>
            <button
              type="submit"
              className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600"
            >
              作成
            </button>
          </form>
        </div>

        <div className="p-4 bg-white border rounded">
          <h2 className="text-xl font-bold mb-4">マイデータ</h2>
          {userData.length === 0 ? (
            <p className="text-gray-500">データがありません</p>
          ) : (
            <ul className="space-y-2">
              {userData.map((item) => (
                <li key={item.id} className="p-2 bg-gray-50 rounded">
                  {item.content}
                  <span className="text-sm text-gray-500 ml-2">
                    {new Date(item.createdAt).toLocaleString()}
                  </span>
                </li>
              ))}
            </ul>
          )}
        </div>

        <div className="mt-8">
          <a href="/" className="text-blue-500 hover:underline">
            ← ホームに戻る
          </a>
        </div>
      </div>
    </main>
  );
}
