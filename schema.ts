import { z } from 'zod';

export const LeadSchema = z.object({
  name: z.string(),
  contact: z.string(),
})
export type Lead = z.infer<typeof LeadSchema>

export const ProjectSchema = z.object({
  project_name: z.string(),
  developer: z.string(),
  tenure: z.string(),
  unit_no: z.string(),
  facing: z.string(),
  layout_type: z.string(),
  area_sqft: z.number(),
  spa_price: z.number(),
})
export type Project = z.infer<typeof ProjectSchema>

export const DiscountSchema = z.object({
  type: z.string(),
  amount: z.number(),
})
export type Discount = z.infer<typeof DiscountSchema>

export const FurnishingSchema = z.object({
  kitchen_cabinet: z.boolean(),
  hood_and_hob: z.boolean(),
  fridge: z.boolean(),
  washing_machine_qty: z.number(),
  airconds_qty: z.number(),
  toilet: z.boolean(),
  heater: z.boolean(),
  shower_screen: z.boolean(),
  wardrobe_qty: z.number(),
  bed_set_qty: z.number(),
  additional_items: z.string().array().nullable(),
})
export type Furnishing = z.infer<typeof FurnishingSchema>

export const OptionSchema = z.object({
  option_name: z.string(),
  rebate: z.number(),
  other_discounts: DiscountSchema.array().nullable(),
  cashback: z.number(),
  down_payment: z.number(),
  nett_price: z.number(),
  loan_amount: z.number(),
  interest_rate: z.number(),
  monthly_instalment: z.number(),
  furnishing: FurnishingSchema,
})
export type Option = z.infer<typeof OptionSchema>

export const LegalFeesSchema = z.object({
  maintenance_fee_psf: z.number(),
  maintenance_fee_total: z.number(),
  included: z.string().array().nullable(),
  not_included: z.string().array().nullable(),
})
export type LegalFees = z.infer<typeof LegalFeesSchema>

export const AgentSchema = z.object({
  name: z.string(),
  phone_number: z.string(),
  email: z.string(),
  signature_url: z.string(),
  logo_url: z.string(),
})
export type Agent = z.infer<typeof AgentSchema>

export const PropertyQuotationSchema = z.object({
  appointment_date: z.coerce.date(),
  quotation_validity: z.coerce.date(),
  lead_info: LeadSchema,
  project_details: ProjectSchema,
  options: OptionSchema.array().nullable(),
  legal_and_fees: LegalFeesSchema,
  agent: AgentSchema,
})
export type PropertyQuotation = z.infer<typeof PropertyQuotationSchema>

;