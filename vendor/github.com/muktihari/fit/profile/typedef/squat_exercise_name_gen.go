// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SquatExerciseName uint16

const (
	SquatExerciseNameLegPress                                        SquatExerciseName = 0
	SquatExerciseNameBackSquatWithBodyBar                            SquatExerciseName = 1
	SquatExerciseNameBackSquats                                      SquatExerciseName = 2
	SquatExerciseNameWeightedBackSquats                              SquatExerciseName = 3
	SquatExerciseNameBalancingSquat                                  SquatExerciseName = 4
	SquatExerciseNameWeightedBalancingSquat                          SquatExerciseName = 5
	SquatExerciseNameBarbellBackSquat                                SquatExerciseName = 6
	SquatExerciseNameBarbellBoxSquat                                 SquatExerciseName = 7
	SquatExerciseNameBarbellFrontSquat                               SquatExerciseName = 8
	SquatExerciseNameBarbellHackSquat                                SquatExerciseName = 9
	SquatExerciseNameBarbellHangSquatSnatch                          SquatExerciseName = 10
	SquatExerciseNameBarbellLateralStepUp                            SquatExerciseName = 11
	SquatExerciseNameBarbellQuarterSquat                             SquatExerciseName = 12
	SquatExerciseNameBarbellSiffSquat                                SquatExerciseName = 13
	SquatExerciseNameBarbellSquatSnatch                              SquatExerciseName = 14
	SquatExerciseNameBarbellSquatWithHeelsRaised                     SquatExerciseName = 15
	SquatExerciseNameBarbellStepover                                 SquatExerciseName = 16
	SquatExerciseNameBarbellStepUp                                   SquatExerciseName = 17
	SquatExerciseNameBenchSquatWithRotationalChop                    SquatExerciseName = 18
	SquatExerciseNameWeightedBenchSquatWithRotationalChop            SquatExerciseName = 19
	SquatExerciseNameBodyWeightWallSquat                             SquatExerciseName = 20
	SquatExerciseNameWeightedWallSquat                               SquatExerciseName = 21
	SquatExerciseNameBoxStepSquat                                    SquatExerciseName = 22
	SquatExerciseNameWeightedBoxStepSquat                            SquatExerciseName = 23
	SquatExerciseNameBracedSquat                                     SquatExerciseName = 24
	SquatExerciseNameCrossedArmBarbellFrontSquat                     SquatExerciseName = 25
	SquatExerciseNameCrossoverDumbbellStepUp                         SquatExerciseName = 26
	SquatExerciseNameDumbbellFrontSquat                              SquatExerciseName = 27
	SquatExerciseNameDumbbellSplitSquat                              SquatExerciseName = 28
	SquatExerciseNameDumbbellSquat                                   SquatExerciseName = 29
	SquatExerciseNameDumbbellSquatClean                              SquatExerciseName = 30
	SquatExerciseNameDumbbellStepover                                SquatExerciseName = 31
	SquatExerciseNameDumbbellStepUp                                  SquatExerciseName = 32
	SquatExerciseNameElevatedSingleLegSquat                          SquatExerciseName = 33
	SquatExerciseNameWeightedElevatedSingleLegSquat                  SquatExerciseName = 34
	SquatExerciseNameFigureFourSquats                                SquatExerciseName = 35
	SquatExerciseNameWeightedFigureFourSquats                        SquatExerciseName = 36
	SquatExerciseNameGobletSquat                                     SquatExerciseName = 37
	SquatExerciseNameKettlebellSquat                                 SquatExerciseName = 38
	SquatExerciseNameKettlebellSwingOverhead                         SquatExerciseName = 39
	SquatExerciseNameKettlebellSwingWithFlipToSquat                  SquatExerciseName = 40
	SquatExerciseNameLateralDumbbellStepUp                           SquatExerciseName = 41
	SquatExerciseNameOneLeggedSquat                                  SquatExerciseName = 42
	SquatExerciseNameOverheadDumbbellSquat                           SquatExerciseName = 43
	SquatExerciseNameOverheadSquat                                   SquatExerciseName = 44
	SquatExerciseNamePartialSingleLegSquat                           SquatExerciseName = 45
	SquatExerciseNameWeightedPartialSingleLegSquat                   SquatExerciseName = 46
	SquatExerciseNamePistolSquat                                     SquatExerciseName = 47
	SquatExerciseNameWeightedPistolSquat                             SquatExerciseName = 48
	SquatExerciseNamePlieSlides                                      SquatExerciseName = 49
	SquatExerciseNameWeightedPlieSlides                              SquatExerciseName = 50
	SquatExerciseNamePlieSquat                                       SquatExerciseName = 51
	SquatExerciseNameWeightedPlieSquat                               SquatExerciseName = 52
	SquatExerciseNamePrisonerSquat                                   SquatExerciseName = 53
	SquatExerciseNameWeightedPrisonerSquat                           SquatExerciseName = 54
	SquatExerciseNameSingleLegBenchGetUp                             SquatExerciseName = 55
	SquatExerciseNameWeightedSingleLegBenchGetUp                     SquatExerciseName = 56
	SquatExerciseNameSingleLegBenchSquat                             SquatExerciseName = 57
	SquatExerciseNameWeightedSingleLegBenchSquat                     SquatExerciseName = 58
	SquatExerciseNameSingleLegSquatOnSwissBall                       SquatExerciseName = 59
	SquatExerciseNameWeightedSingleLegSquatOnSwissBall               SquatExerciseName = 60
	SquatExerciseNameSquat                                           SquatExerciseName = 61
	SquatExerciseNameWeightedSquat                                   SquatExerciseName = 62
	SquatExerciseNameSquatsWithBand                                  SquatExerciseName = 63
	SquatExerciseNameStaggeredSquat                                  SquatExerciseName = 64
	SquatExerciseNameWeightedStaggeredSquat                          SquatExerciseName = 65
	SquatExerciseNameStepUp                                          SquatExerciseName = 66
	SquatExerciseNameWeightedStepUp                                  SquatExerciseName = 67
	SquatExerciseNameSuitcaseSquats                                  SquatExerciseName = 68
	SquatExerciseNameSumoSquat                                       SquatExerciseName = 69
	SquatExerciseNameSumoSquatSlideIn                                SquatExerciseName = 70
	SquatExerciseNameWeightedSumoSquatSlideIn                        SquatExerciseName = 71
	SquatExerciseNameSumoSquatToHighPull                             SquatExerciseName = 72
	SquatExerciseNameSumoSquatToStand                                SquatExerciseName = 73
	SquatExerciseNameWeightedSumoSquatToStand                        SquatExerciseName = 74
	SquatExerciseNameSumoSquatWithRotation                           SquatExerciseName = 75
	SquatExerciseNameWeightedSumoSquatWithRotation                   SquatExerciseName = 76
	SquatExerciseNameSwissBallBodyWeightWallSquat                    SquatExerciseName = 77
	SquatExerciseNameWeightedSwissBallWallSquat                      SquatExerciseName = 78
	SquatExerciseNameThrusters                                       SquatExerciseName = 79
	SquatExerciseNameUnevenSquat                                     SquatExerciseName = 80
	SquatExerciseNameWeightedUnevenSquat                             SquatExerciseName = 81
	SquatExerciseNameWaistSlimmingSquat                              SquatExerciseName = 82
	SquatExerciseNameWallBall                                        SquatExerciseName = 83
	SquatExerciseNameWideStanceBarbellSquat                          SquatExerciseName = 84
	SquatExerciseNameWideStanceGobletSquat                           SquatExerciseName = 85
	SquatExerciseNameZercherSquat                                    SquatExerciseName = 86
	SquatExerciseNameKbsOverhead                                     SquatExerciseName = 87 // Deprecated do not use
	SquatExerciseNameSquatAndSideKick                                SquatExerciseName = 88
	SquatExerciseNameSquatJumpsInNOut                                SquatExerciseName = 89
	SquatExerciseNamePilatesPlieSquatsParallelTurnedOutFlatAndHeels  SquatExerciseName = 90
	SquatExerciseNameReleveStraightLegAndKneeBentWithOneLegVariation SquatExerciseName = 91
	SquatExerciseNameInvalid                                         SquatExerciseName = 0xFFFF
)

func (s SquatExerciseName) Uint16() uint16 { return uint16(s) }

func (s SquatExerciseName) String() string {
	switch s {
	case SquatExerciseNameLegPress:
		return "leg_press"
	case SquatExerciseNameBackSquatWithBodyBar:
		return "back_squat_with_body_bar"
	case SquatExerciseNameBackSquats:
		return "back_squats"
	case SquatExerciseNameWeightedBackSquats:
		return "weighted_back_squats"
	case SquatExerciseNameBalancingSquat:
		return "balancing_squat"
	case SquatExerciseNameWeightedBalancingSquat:
		return "weighted_balancing_squat"
	case SquatExerciseNameBarbellBackSquat:
		return "barbell_back_squat"
	case SquatExerciseNameBarbellBoxSquat:
		return "barbell_box_squat"
	case SquatExerciseNameBarbellFrontSquat:
		return "barbell_front_squat"
	case SquatExerciseNameBarbellHackSquat:
		return "barbell_hack_squat"
	case SquatExerciseNameBarbellHangSquatSnatch:
		return "barbell_hang_squat_snatch"
	case SquatExerciseNameBarbellLateralStepUp:
		return "barbell_lateral_step_up"
	case SquatExerciseNameBarbellQuarterSquat:
		return "barbell_quarter_squat"
	case SquatExerciseNameBarbellSiffSquat:
		return "barbell_siff_squat"
	case SquatExerciseNameBarbellSquatSnatch:
		return "barbell_squat_snatch"
	case SquatExerciseNameBarbellSquatWithHeelsRaised:
		return "barbell_squat_with_heels_raised"
	case SquatExerciseNameBarbellStepover:
		return "barbell_stepover"
	case SquatExerciseNameBarbellStepUp:
		return "barbell_step_up"
	case SquatExerciseNameBenchSquatWithRotationalChop:
		return "bench_squat_with_rotational_chop"
	case SquatExerciseNameWeightedBenchSquatWithRotationalChop:
		return "weighted_bench_squat_with_rotational_chop"
	case SquatExerciseNameBodyWeightWallSquat:
		return "body_weight_wall_squat"
	case SquatExerciseNameWeightedWallSquat:
		return "weighted_wall_squat"
	case SquatExerciseNameBoxStepSquat:
		return "box_step_squat"
	case SquatExerciseNameWeightedBoxStepSquat:
		return "weighted_box_step_squat"
	case SquatExerciseNameBracedSquat:
		return "braced_squat"
	case SquatExerciseNameCrossedArmBarbellFrontSquat:
		return "crossed_arm_barbell_front_squat"
	case SquatExerciseNameCrossoverDumbbellStepUp:
		return "crossover_dumbbell_step_up"
	case SquatExerciseNameDumbbellFrontSquat:
		return "dumbbell_front_squat"
	case SquatExerciseNameDumbbellSplitSquat:
		return "dumbbell_split_squat"
	case SquatExerciseNameDumbbellSquat:
		return "dumbbell_squat"
	case SquatExerciseNameDumbbellSquatClean:
		return "dumbbell_squat_clean"
	case SquatExerciseNameDumbbellStepover:
		return "dumbbell_stepover"
	case SquatExerciseNameDumbbellStepUp:
		return "dumbbell_step_up"
	case SquatExerciseNameElevatedSingleLegSquat:
		return "elevated_single_leg_squat"
	case SquatExerciseNameWeightedElevatedSingleLegSquat:
		return "weighted_elevated_single_leg_squat"
	case SquatExerciseNameFigureFourSquats:
		return "figure_four_squats"
	case SquatExerciseNameWeightedFigureFourSquats:
		return "weighted_figure_four_squats"
	case SquatExerciseNameGobletSquat:
		return "goblet_squat"
	case SquatExerciseNameKettlebellSquat:
		return "kettlebell_squat"
	case SquatExerciseNameKettlebellSwingOverhead:
		return "kettlebell_swing_overhead"
	case SquatExerciseNameKettlebellSwingWithFlipToSquat:
		return "kettlebell_swing_with_flip_to_squat"
	case SquatExerciseNameLateralDumbbellStepUp:
		return "lateral_dumbbell_step_up"
	case SquatExerciseNameOneLeggedSquat:
		return "one_legged_squat"
	case SquatExerciseNameOverheadDumbbellSquat:
		return "overhead_dumbbell_squat"
	case SquatExerciseNameOverheadSquat:
		return "overhead_squat"
	case SquatExerciseNamePartialSingleLegSquat:
		return "partial_single_leg_squat"
	case SquatExerciseNameWeightedPartialSingleLegSquat:
		return "weighted_partial_single_leg_squat"
	case SquatExerciseNamePistolSquat:
		return "pistol_squat"
	case SquatExerciseNameWeightedPistolSquat:
		return "weighted_pistol_squat"
	case SquatExerciseNamePlieSlides:
		return "plie_slides"
	case SquatExerciseNameWeightedPlieSlides:
		return "weighted_plie_slides"
	case SquatExerciseNamePlieSquat:
		return "plie_squat"
	case SquatExerciseNameWeightedPlieSquat:
		return "weighted_plie_squat"
	case SquatExerciseNamePrisonerSquat:
		return "prisoner_squat"
	case SquatExerciseNameWeightedPrisonerSquat:
		return "weighted_prisoner_squat"
	case SquatExerciseNameSingleLegBenchGetUp:
		return "single_leg_bench_get_up"
	case SquatExerciseNameWeightedSingleLegBenchGetUp:
		return "weighted_single_leg_bench_get_up"
	case SquatExerciseNameSingleLegBenchSquat:
		return "single_leg_bench_squat"
	case SquatExerciseNameWeightedSingleLegBenchSquat:
		return "weighted_single_leg_bench_squat"
	case SquatExerciseNameSingleLegSquatOnSwissBall:
		return "single_leg_squat_on_swiss_ball"
	case SquatExerciseNameWeightedSingleLegSquatOnSwissBall:
		return "weighted_single_leg_squat_on_swiss_ball"
	case SquatExerciseNameSquat:
		return "squat"
	case SquatExerciseNameWeightedSquat:
		return "weighted_squat"
	case SquatExerciseNameSquatsWithBand:
		return "squats_with_band"
	case SquatExerciseNameStaggeredSquat:
		return "staggered_squat"
	case SquatExerciseNameWeightedStaggeredSquat:
		return "weighted_staggered_squat"
	case SquatExerciseNameStepUp:
		return "step_up"
	case SquatExerciseNameWeightedStepUp:
		return "weighted_step_up"
	case SquatExerciseNameSuitcaseSquats:
		return "suitcase_squats"
	case SquatExerciseNameSumoSquat:
		return "sumo_squat"
	case SquatExerciseNameSumoSquatSlideIn:
		return "sumo_squat_slide_in"
	case SquatExerciseNameWeightedSumoSquatSlideIn:
		return "weighted_sumo_squat_slide_in"
	case SquatExerciseNameSumoSquatToHighPull:
		return "sumo_squat_to_high_pull"
	case SquatExerciseNameSumoSquatToStand:
		return "sumo_squat_to_stand"
	case SquatExerciseNameWeightedSumoSquatToStand:
		return "weighted_sumo_squat_to_stand"
	case SquatExerciseNameSumoSquatWithRotation:
		return "sumo_squat_with_rotation"
	case SquatExerciseNameWeightedSumoSquatWithRotation:
		return "weighted_sumo_squat_with_rotation"
	case SquatExerciseNameSwissBallBodyWeightWallSquat:
		return "swiss_ball_body_weight_wall_squat"
	case SquatExerciseNameWeightedSwissBallWallSquat:
		return "weighted_swiss_ball_wall_squat"
	case SquatExerciseNameThrusters:
		return "thrusters"
	case SquatExerciseNameUnevenSquat:
		return "uneven_squat"
	case SquatExerciseNameWeightedUnevenSquat:
		return "weighted_uneven_squat"
	case SquatExerciseNameWaistSlimmingSquat:
		return "waist_slimming_squat"
	case SquatExerciseNameWallBall:
		return "wall_ball"
	case SquatExerciseNameWideStanceBarbellSquat:
		return "wide_stance_barbell_squat"
	case SquatExerciseNameWideStanceGobletSquat:
		return "wide_stance_goblet_squat"
	case SquatExerciseNameZercherSquat:
		return "zercher_squat"
	case SquatExerciseNameKbsOverhead:
		return "kbs_overhead"
	case SquatExerciseNameSquatAndSideKick:
		return "squat_and_side_kick"
	case SquatExerciseNameSquatJumpsInNOut:
		return "squat_jumps_in_n_out"
	case SquatExerciseNamePilatesPlieSquatsParallelTurnedOutFlatAndHeels:
		return "pilates_plie_squats_parallel_turned_out_flat_and_heels"
	case SquatExerciseNameReleveStraightLegAndKneeBentWithOneLegVariation:
		return "releve_straight_leg_and_knee_bent_with_one_leg_variation"
	default:
		return "SquatExerciseNameInvalid(" + strconv.FormatUint(uint64(s), 10) + ")"
	}
}

// FromString parse string into SquatExerciseName constant it's represent, return SquatExerciseNameInvalid if not found.
func SquatExerciseNameFromString(s string) SquatExerciseName {
	switch s {
	case "leg_press":
		return SquatExerciseNameLegPress
	case "back_squat_with_body_bar":
		return SquatExerciseNameBackSquatWithBodyBar
	case "back_squats":
		return SquatExerciseNameBackSquats
	case "weighted_back_squats":
		return SquatExerciseNameWeightedBackSquats
	case "balancing_squat":
		return SquatExerciseNameBalancingSquat
	case "weighted_balancing_squat":
		return SquatExerciseNameWeightedBalancingSquat
	case "barbell_back_squat":
		return SquatExerciseNameBarbellBackSquat
	case "barbell_box_squat":
		return SquatExerciseNameBarbellBoxSquat
	case "barbell_front_squat":
		return SquatExerciseNameBarbellFrontSquat
	case "barbell_hack_squat":
		return SquatExerciseNameBarbellHackSquat
	case "barbell_hang_squat_snatch":
		return SquatExerciseNameBarbellHangSquatSnatch
	case "barbell_lateral_step_up":
		return SquatExerciseNameBarbellLateralStepUp
	case "barbell_quarter_squat":
		return SquatExerciseNameBarbellQuarterSquat
	case "barbell_siff_squat":
		return SquatExerciseNameBarbellSiffSquat
	case "barbell_squat_snatch":
		return SquatExerciseNameBarbellSquatSnatch
	case "barbell_squat_with_heels_raised":
		return SquatExerciseNameBarbellSquatWithHeelsRaised
	case "barbell_stepover":
		return SquatExerciseNameBarbellStepover
	case "barbell_step_up":
		return SquatExerciseNameBarbellStepUp
	case "bench_squat_with_rotational_chop":
		return SquatExerciseNameBenchSquatWithRotationalChop
	case "weighted_bench_squat_with_rotational_chop":
		return SquatExerciseNameWeightedBenchSquatWithRotationalChop
	case "body_weight_wall_squat":
		return SquatExerciseNameBodyWeightWallSquat
	case "weighted_wall_squat":
		return SquatExerciseNameWeightedWallSquat
	case "box_step_squat":
		return SquatExerciseNameBoxStepSquat
	case "weighted_box_step_squat":
		return SquatExerciseNameWeightedBoxStepSquat
	case "braced_squat":
		return SquatExerciseNameBracedSquat
	case "crossed_arm_barbell_front_squat":
		return SquatExerciseNameCrossedArmBarbellFrontSquat
	case "crossover_dumbbell_step_up":
		return SquatExerciseNameCrossoverDumbbellStepUp
	case "dumbbell_front_squat":
		return SquatExerciseNameDumbbellFrontSquat
	case "dumbbell_split_squat":
		return SquatExerciseNameDumbbellSplitSquat
	case "dumbbell_squat":
		return SquatExerciseNameDumbbellSquat
	case "dumbbell_squat_clean":
		return SquatExerciseNameDumbbellSquatClean
	case "dumbbell_stepover":
		return SquatExerciseNameDumbbellStepover
	case "dumbbell_step_up":
		return SquatExerciseNameDumbbellStepUp
	case "elevated_single_leg_squat":
		return SquatExerciseNameElevatedSingleLegSquat
	case "weighted_elevated_single_leg_squat":
		return SquatExerciseNameWeightedElevatedSingleLegSquat
	case "figure_four_squats":
		return SquatExerciseNameFigureFourSquats
	case "weighted_figure_four_squats":
		return SquatExerciseNameWeightedFigureFourSquats
	case "goblet_squat":
		return SquatExerciseNameGobletSquat
	case "kettlebell_squat":
		return SquatExerciseNameKettlebellSquat
	case "kettlebell_swing_overhead":
		return SquatExerciseNameKettlebellSwingOverhead
	case "kettlebell_swing_with_flip_to_squat":
		return SquatExerciseNameKettlebellSwingWithFlipToSquat
	case "lateral_dumbbell_step_up":
		return SquatExerciseNameLateralDumbbellStepUp
	case "one_legged_squat":
		return SquatExerciseNameOneLeggedSquat
	case "overhead_dumbbell_squat":
		return SquatExerciseNameOverheadDumbbellSquat
	case "overhead_squat":
		return SquatExerciseNameOverheadSquat
	case "partial_single_leg_squat":
		return SquatExerciseNamePartialSingleLegSquat
	case "weighted_partial_single_leg_squat":
		return SquatExerciseNameWeightedPartialSingleLegSquat
	case "pistol_squat":
		return SquatExerciseNamePistolSquat
	case "weighted_pistol_squat":
		return SquatExerciseNameWeightedPistolSquat
	case "plie_slides":
		return SquatExerciseNamePlieSlides
	case "weighted_plie_slides":
		return SquatExerciseNameWeightedPlieSlides
	case "plie_squat":
		return SquatExerciseNamePlieSquat
	case "weighted_plie_squat":
		return SquatExerciseNameWeightedPlieSquat
	case "prisoner_squat":
		return SquatExerciseNamePrisonerSquat
	case "weighted_prisoner_squat":
		return SquatExerciseNameWeightedPrisonerSquat
	case "single_leg_bench_get_up":
		return SquatExerciseNameSingleLegBenchGetUp
	case "weighted_single_leg_bench_get_up":
		return SquatExerciseNameWeightedSingleLegBenchGetUp
	case "single_leg_bench_squat":
		return SquatExerciseNameSingleLegBenchSquat
	case "weighted_single_leg_bench_squat":
		return SquatExerciseNameWeightedSingleLegBenchSquat
	case "single_leg_squat_on_swiss_ball":
		return SquatExerciseNameSingleLegSquatOnSwissBall
	case "weighted_single_leg_squat_on_swiss_ball":
		return SquatExerciseNameWeightedSingleLegSquatOnSwissBall
	case "squat":
		return SquatExerciseNameSquat
	case "weighted_squat":
		return SquatExerciseNameWeightedSquat
	case "squats_with_band":
		return SquatExerciseNameSquatsWithBand
	case "staggered_squat":
		return SquatExerciseNameStaggeredSquat
	case "weighted_staggered_squat":
		return SquatExerciseNameWeightedStaggeredSquat
	case "step_up":
		return SquatExerciseNameStepUp
	case "weighted_step_up":
		return SquatExerciseNameWeightedStepUp
	case "suitcase_squats":
		return SquatExerciseNameSuitcaseSquats
	case "sumo_squat":
		return SquatExerciseNameSumoSquat
	case "sumo_squat_slide_in":
		return SquatExerciseNameSumoSquatSlideIn
	case "weighted_sumo_squat_slide_in":
		return SquatExerciseNameWeightedSumoSquatSlideIn
	case "sumo_squat_to_high_pull":
		return SquatExerciseNameSumoSquatToHighPull
	case "sumo_squat_to_stand":
		return SquatExerciseNameSumoSquatToStand
	case "weighted_sumo_squat_to_stand":
		return SquatExerciseNameWeightedSumoSquatToStand
	case "sumo_squat_with_rotation":
		return SquatExerciseNameSumoSquatWithRotation
	case "weighted_sumo_squat_with_rotation":
		return SquatExerciseNameWeightedSumoSquatWithRotation
	case "swiss_ball_body_weight_wall_squat":
		return SquatExerciseNameSwissBallBodyWeightWallSquat
	case "weighted_swiss_ball_wall_squat":
		return SquatExerciseNameWeightedSwissBallWallSquat
	case "thrusters":
		return SquatExerciseNameThrusters
	case "uneven_squat":
		return SquatExerciseNameUnevenSquat
	case "weighted_uneven_squat":
		return SquatExerciseNameWeightedUnevenSquat
	case "waist_slimming_squat":
		return SquatExerciseNameWaistSlimmingSquat
	case "wall_ball":
		return SquatExerciseNameWallBall
	case "wide_stance_barbell_squat":
		return SquatExerciseNameWideStanceBarbellSquat
	case "wide_stance_goblet_squat":
		return SquatExerciseNameWideStanceGobletSquat
	case "zercher_squat":
		return SquatExerciseNameZercherSquat
	case "kbs_overhead":
		return SquatExerciseNameKbsOverhead
	case "squat_and_side_kick":
		return SquatExerciseNameSquatAndSideKick
	case "squat_jumps_in_n_out":
		return SquatExerciseNameSquatJumpsInNOut
	case "pilates_plie_squats_parallel_turned_out_flat_and_heels":
		return SquatExerciseNamePilatesPlieSquatsParallelTurnedOutFlatAndHeels
	case "releve_straight_leg_and_knee_bent_with_one_leg_variation":
		return SquatExerciseNameReleveStraightLegAndKneeBentWithOneLegVariation
	default:
		return SquatExerciseNameInvalid
	}
}

// List returns all constants.
func ListSquatExerciseName() []SquatExerciseName {
	return []SquatExerciseName{
		SquatExerciseNameLegPress,
		SquatExerciseNameBackSquatWithBodyBar,
		SquatExerciseNameBackSquats,
		SquatExerciseNameWeightedBackSquats,
		SquatExerciseNameBalancingSquat,
		SquatExerciseNameWeightedBalancingSquat,
		SquatExerciseNameBarbellBackSquat,
		SquatExerciseNameBarbellBoxSquat,
		SquatExerciseNameBarbellFrontSquat,
		SquatExerciseNameBarbellHackSquat,
		SquatExerciseNameBarbellHangSquatSnatch,
		SquatExerciseNameBarbellLateralStepUp,
		SquatExerciseNameBarbellQuarterSquat,
		SquatExerciseNameBarbellSiffSquat,
		SquatExerciseNameBarbellSquatSnatch,
		SquatExerciseNameBarbellSquatWithHeelsRaised,
		SquatExerciseNameBarbellStepover,
		SquatExerciseNameBarbellStepUp,
		SquatExerciseNameBenchSquatWithRotationalChop,
		SquatExerciseNameWeightedBenchSquatWithRotationalChop,
		SquatExerciseNameBodyWeightWallSquat,
		SquatExerciseNameWeightedWallSquat,
		SquatExerciseNameBoxStepSquat,
		SquatExerciseNameWeightedBoxStepSquat,
		SquatExerciseNameBracedSquat,
		SquatExerciseNameCrossedArmBarbellFrontSquat,
		SquatExerciseNameCrossoverDumbbellStepUp,
		SquatExerciseNameDumbbellFrontSquat,
		SquatExerciseNameDumbbellSplitSquat,
		SquatExerciseNameDumbbellSquat,
		SquatExerciseNameDumbbellSquatClean,
		SquatExerciseNameDumbbellStepover,
		SquatExerciseNameDumbbellStepUp,
		SquatExerciseNameElevatedSingleLegSquat,
		SquatExerciseNameWeightedElevatedSingleLegSquat,
		SquatExerciseNameFigureFourSquats,
		SquatExerciseNameWeightedFigureFourSquats,
		SquatExerciseNameGobletSquat,
		SquatExerciseNameKettlebellSquat,
		SquatExerciseNameKettlebellSwingOverhead,
		SquatExerciseNameKettlebellSwingWithFlipToSquat,
		SquatExerciseNameLateralDumbbellStepUp,
		SquatExerciseNameOneLeggedSquat,
		SquatExerciseNameOverheadDumbbellSquat,
		SquatExerciseNameOverheadSquat,
		SquatExerciseNamePartialSingleLegSquat,
		SquatExerciseNameWeightedPartialSingleLegSquat,
		SquatExerciseNamePistolSquat,
		SquatExerciseNameWeightedPistolSquat,
		SquatExerciseNamePlieSlides,
		SquatExerciseNameWeightedPlieSlides,
		SquatExerciseNamePlieSquat,
		SquatExerciseNameWeightedPlieSquat,
		SquatExerciseNamePrisonerSquat,
		SquatExerciseNameWeightedPrisonerSquat,
		SquatExerciseNameSingleLegBenchGetUp,
		SquatExerciseNameWeightedSingleLegBenchGetUp,
		SquatExerciseNameSingleLegBenchSquat,
		SquatExerciseNameWeightedSingleLegBenchSquat,
		SquatExerciseNameSingleLegSquatOnSwissBall,
		SquatExerciseNameWeightedSingleLegSquatOnSwissBall,
		SquatExerciseNameSquat,
		SquatExerciseNameWeightedSquat,
		SquatExerciseNameSquatsWithBand,
		SquatExerciseNameStaggeredSquat,
		SquatExerciseNameWeightedStaggeredSquat,
		SquatExerciseNameStepUp,
		SquatExerciseNameWeightedStepUp,
		SquatExerciseNameSuitcaseSquats,
		SquatExerciseNameSumoSquat,
		SquatExerciseNameSumoSquatSlideIn,
		SquatExerciseNameWeightedSumoSquatSlideIn,
		SquatExerciseNameSumoSquatToHighPull,
		SquatExerciseNameSumoSquatToStand,
		SquatExerciseNameWeightedSumoSquatToStand,
		SquatExerciseNameSumoSquatWithRotation,
		SquatExerciseNameWeightedSumoSquatWithRotation,
		SquatExerciseNameSwissBallBodyWeightWallSquat,
		SquatExerciseNameWeightedSwissBallWallSquat,
		SquatExerciseNameThrusters,
		SquatExerciseNameUnevenSquat,
		SquatExerciseNameWeightedUnevenSquat,
		SquatExerciseNameWaistSlimmingSquat,
		SquatExerciseNameWallBall,
		SquatExerciseNameWideStanceBarbellSquat,
		SquatExerciseNameWideStanceGobletSquat,
		SquatExerciseNameZercherSquat,
		SquatExerciseNameKbsOverhead,
		SquatExerciseNameSquatAndSideKick,
		SquatExerciseNameSquatJumpsInNOut,
		SquatExerciseNamePilatesPlieSquatsParallelTurnedOutFlatAndHeels,
		SquatExerciseNameReleveStraightLegAndKneeBentWithOneLegVariation,
	}
}