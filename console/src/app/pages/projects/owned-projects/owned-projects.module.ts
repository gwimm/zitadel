import { CommonModule } from '@angular/common';
import { NgModule, NO_ERRORS_SCHEMA } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatChipsModule } from '@angular/material/chips';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSelectModule } from '@angular/material/select';
import { MatSortModule } from '@angular/material/sort';
import { MatTableModule } from '@angular/material/table';
import { MatTabsModule } from '@angular/material/tabs';
import { MatTooltipModule } from '@angular/material/tooltip';
import { TranslateModule } from '@ngx-translate/core';
import { HasRoleModule } from 'src/app/directives/has-role/has-role.module';
import { AvatarModule } from 'src/app/modules/avatar/avatar.module';
import { CardModule } from 'src/app/modules/card/card.module';
import { ChangesModule } from 'src/app/modules/changes/changes.module';
import { MetaLayoutModule } from 'src/app/modules/meta-layout/meta-layout.module';
import { ProjectContributorsModule } from 'src/app/modules/project-contributors/project-contributors.module';
import { ProjectRolesModule } from 'src/app/modules/project-roles/project-roles.module';
import { UserGrantsModule } from 'src/app/modules/user-grants/user-grants.module';
import { WarnDialogModule } from 'src/app/modules/warn-dialog/warn-dialog.module';
import { PipesModule } from 'src/app/pipes/pipes.module';

import { OwnedProjectDetailComponent } from './owned-project-detail/owned-project-detail.component';
import { OwnedProjectGridComponent } from './owned-project-grid/owned-project-grid.component';
import { OwnedProjectListComponent } from './owned-project-list/owned-project-list.component';
import { OwnedProjectsRoutingModule } from './owned-projects-routing.module';
import { OwnedProjectsComponent } from './owned-projects.component';
import { ProjectApplicationGridComponent } from './project-application-grid/project-application-grid.component';
import { ProjectApplicationsComponent } from './project-applications/project-applications.component';
import { ProjectGrantsComponent } from './project-grants/project-grants.component';

@NgModule({
    declarations: [
        OwnedProjectsComponent,
        OwnedProjectListComponent,
        OwnedProjectGridComponent,
        OwnedProjectDetailComponent,
        ProjectApplicationGridComponent,
        ProjectApplicationsComponent,
        ProjectGrantsComponent,
    ],
    imports: [
        CommonModule,
        OwnedProjectsRoutingModule,
        UserGrantsModule,
        ProjectContributorsModule,
        FormsModule,
        ReactiveFormsModule,
        TranslateModule,
        AvatarModule,
        ReactiveFormsModule,
        HasRoleModule,
        MatTableModule,
        MatPaginatorModule,
        MatFormFieldModule,
        MatInputModule,
        ChangesModule,
        MatChipsModule,
        MatIconModule,
        MatButtonModule,
        WarnDialogModule,
        MatProgressSpinnerModule,
        MetaLayoutModule,
        MatProgressBarModule,
        ProjectRolesModule,
        MatTabsModule,
        MatCheckboxModule,
        CardModule,
        MatSelectModule,
        MatTooltipModule,
        MatSortModule,
        PipesModule,
        TranslateModule,
    ],
    schemas: [NO_ERRORS_SCHEMA],
})
export class OwnedProjectsModule { }