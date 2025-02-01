import type { Config, ConfigId } from "@/types/config";
import axios from "axios";

class ConfigService {
    private baseUrl: string;

    constructor() {
        this.baseUrl = "http://localhost:8080/configs"
    }

    async getAll(): Promise<Config[]> {
        return (await axios.get<Config[]>(`${this.baseUrl}/all`)).data
    }

    async save(dto: Config): Promise<void> {
        return axios.put(`${this.baseUrl}/save`, dto)
    }

    async del(id: ConfigId): Promise<void> {
        return axios.delete(`${this.baseUrl}/${id}`)
    }
}

export const configService = new ConfigService;