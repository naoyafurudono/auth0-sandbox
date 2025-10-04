export class ApiClient {
  private baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  async getUserProfile() {
    const response = await fetch(`${this.baseUrl}/api/v1/users/me/profile`, {
      credentials: "include",
    });

    if (response.status === 404) {
      return null;
    }

    if (!response.ok) {
      throw new Error("Failed to fetch profile");
    }

    return response.json();
  }

  async updateUserProfile(data: {
    displayName?: string;
    bio?: string;
    avatarUrl?: string;
  }) {
    const response = await fetch(`${this.baseUrl}/api/v1/users/me/profile`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      throw new Error("Failed to update profile");
    }

    return response.json();
  }

  async getUserData() {
    const response = await fetch(`${this.baseUrl}/api/v1/users/me/data`, {
      credentials: "include",
    });

    if (!response.ok) {
      throw new Error("Failed to fetch user data");
    }

    return response.json();
  }

  async createUserData(content: string) {
    const response = await fetch(`${this.baseUrl}/api/v1/users/me/data`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify({ content }),
    });

    if (!response.ok) {
      throw new Error("Failed to create user data");
    }

    return response.json();
  }
}

export const apiClient = new ApiClient(
  process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080"
);
