package permission

import (
	"github.com/OpenSlides/openslides-permission-service/internal/collection"
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

func openSlidesCollections(dp dataprovider.DataProvider) []perm.Connecter {
	return []perm.Connecter{
		collection.AgendaItem(dp),
		collection.ListOfSpeaker(dp),
		collection.Assignment(dp),
		collection.Mediafile(dp),
		collection.Motion(dp),
		collection.Poll(dp),
		collection.PersonalNote(dp),
		collection.User(dp),
		collection.Meeting(dp),
		collection.Committee(dp),

		collection.Public(dp, "resource", "organisation"),
		collection.ReadInMeeting(dp, "tag", "group"),
		collection.ReadPerm(dp, perm.AssignmentCanSee, "assignment", "assignment_candidate"),
		collection.ReadPerm(dp, perm.AgendaItemCanSee, "topic"),
		collection.ReadPerm(
			dp,
			perm.ProjectorCanSee,
			"projector",
			"projection",
			"projectiondefault",
			"projector_message",
			"projector_countdown",
		),
		collection.ReadPerm(
			dp,
			perm.MotionCanSee,
			"motion_workflow",
			"motion_category",
			"motion_state",
			"motion_statute_paragraph",
		),

		collection.OrgaManager(
			dp,
			"resource.delete",
			"organisation.update",
			"committee.create",
			"committee.update",
			"committee.delete",
		),

		collection.WritePerm(dp, map[string]perm.TPermission{
			"agenda_item.assign":                       perm.AgendaItemCanManage,
			"agenda_item.create":                       perm.AgendaItemCanManage,
			"agenda_item.delete":                       perm.AgendaItemCanManage,
			"agenda_item.numbering":                    perm.AgendaItemCanManage,
			"agenda_item.sort":                         perm.AgendaItemCanManage,
			"agenda_item.update":                       perm.AgendaItemCanManage,
			"assignment.create":                        perm.AssignmentCanManage,
			"assignment.delete":                        perm.AssignmentCanManage,
			"assignment.update":                        perm.AssignmentCanManage,
			"group.create":                             perm.UserCanManage,
			"group.delete":                             perm.UserCanManage,
			"group.set_permission":                     perm.UserCanManage,
			"group.update":                             perm.UserCanManage,
			"list_of_speakers.delete_all_speakers":     perm.ListOfSpeakersCanManage,
			"list_of_speakers.re_add_last":             perm.ListOfSpeakersCanManage,
			"list_of_speakers.update":                  perm.ListOfSpeakersCanManage,
			"mediafile.create_directory":               perm.MediafileCanManage,
			"mediafile.delete":                         perm.MediafileCanManage,
			"mediafile.move":                           perm.MediafileCanManage,
			"mediafile.update":                         perm.MediafileCanManage,
			"mediafile.upload":                         perm.MediafileCanManage,
			"meeting.delete_all_speakers_of_all_lists": perm.ListOfSpeakersCanManage,
			"meeting.set_font":                         perm.MeetingCanManageLogosAndFonts,
			"meeting.set_logo":                         perm.MeetingCanManageLogosAndFonts,
			"meeting.unset_font":                       perm.MeetingCanManageLogosAndFonts,
			"meeting.unset_logo":                       perm.MeetingCanManageLogosAndFonts,
			"motion.update_metadata":                   perm.MotionCanManageMetadata,
			"motion.follow_recommendation":             perm.MotionCanManageMetadata,
			"motion.reset_recommendation":              perm.MotionCanManageMetadata,
			"motion.reset_state":                       perm.MotionCanManageMetadata,
			"motion.set_recommendation":                perm.MotionCanManageMetadata,
			"motion.sort":                              perm.MotionCanManageMetadata,
			"motion_block.create":                      perm.MotionCanManage,
			"motion_block.delete":                      perm.MotionCanManage,
			"motion_block.update":                      perm.MotionCanManage,
			"motion_category.create":                   perm.MotionCanManage,
			"motion_category.delete":                   perm.MotionCanManage,
			"motion_category.number_motions":           perm.MotionCanManage,
			"motion_category.sort":                     perm.MotionCanManage,
			"motion_category.sort_motions_in_category": perm.MotionCanManage,
			"motion_category.update":                   perm.MotionCanManage,
			"motion_change_recommendation.create":      perm.MotionCanManage,
			"motion_change_recommendation.delete":      perm.MotionCanManage,
			"motion_change_recommendation.update":      perm.MotionCanManage,
			"motion_comment_section.create":            perm.MotionCanManage,
			"motion_comment_section.delete":            perm.MotionCanManage,
			"motion_comment_section.sort":              perm.MotionCanManage,
			"motion_comment_section.update":            perm.MotionCanManage,
			"motion_state.create":                      perm.MotionCanManage,
			"motion_state.delete":                      perm.MotionCanManage,
			"motion_state.update":                      perm.MotionCanManage,
			"motion_statute_paragraph.create":          perm.MotionCanManage,
			"motion_statute_paragraph.delete":          perm.MotionCanManage,
			"motion_statute_paragraph.sort":            perm.MotionCanManage,
			"motion_statute_paragraph.update":          perm.MotionCanManage,
			"motion_submitter.delete":                  perm.MotionCanManage,
			"motion_submitter.sort":                    perm.MotionCanManage,
			"motion_workflow.create":                   perm.MotionCanManage,
			"motion_workflow.delete":                   perm.MotionCanManage,
			"motion_workflow.update":                   perm.MotionCanManage,
			"speaker.end_speech":                       perm.ListOfSpeakersCanManage,
			"speaker.sort":                             perm.ListOfSpeakersCanManage,
			"speaker.speak":                            perm.ListOfSpeakersCanManage,
			"speaker.update":                           perm.ListOfSpeakersCanManage,
			"tag.create":                               perm.TagCanManage,
			"tag.delete":                               perm.TagCanManage,
			"tag.update":                               perm.TagCanManage,
			"topic.create":                             perm.AgendaItemCanManage,
			"topic.delete":                             perm.AgendaItemCanManage,
			"topic.update":                             perm.AgendaItemCanManage,
			"user.create_temporary":                    perm.UserCanManage,
			"user.delete_temporary":                    perm.UserCanManage,
			"user.generate_new_password_temporary":     perm.UserCanManage,
			"user.reset_password_to_default_temporary": perm.UserCanManage,
			"user.update_temporary":                    perm.UserCanManage,
		}),
	}
}
