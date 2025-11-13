// Sample TypeScript file for testing

export function greet(name: string): string {
  return `Hello, ${name}!`;
}

function calculateSum(a: number, b: number): number {
  return a + b;
}

export async function fetchData(url: string): Promise<any> {
  const response = await fetch(url);
  return response.json();
}

class UserService {
  constructor(private apiUrl: string) {}

  async getUser(id: number) {
    return fetchData(`${this.apiUrl}/users/${id}`);
  }

  updateUser(id: number, data: any) {
    console.log('Updating user', id, data);
  }
}

const arrowFunction = (x: number) => x * 2;

export default UserService;
